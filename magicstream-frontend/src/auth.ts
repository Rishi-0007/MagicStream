import NextAuth, { type DefaultSession, type NextAuthOptions } from "next-auth";
import GoogleProvider from "next-auth/providers/google";
import CredentialsProvider from "next-auth/providers/credentials";

declare module "next-auth" {
  interface Session extends DefaultSession {
    backendAccess?: string;
    backendRefresh?: string;
  }
}

async function loginWithBackend(email: string, password: string) {
  const base = process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080";
  const res = await fetch(`${base}/api/auth/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });
  if (!res.ok) return null;
  return res.json() as Promise<{
    access_token: string;
    refresh_token: string;
    user: { email: string; name: string; avatar_url?: string };
  }>;
}

async function exchangeGoogleIdToken(idToken: string) {
  const base = process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080";
  const res = await fetch(`${base}/api/auth/oauth/google`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ id_token: idToken }),
  });
  if (!res.ok) return null;
  return res.json() as Promise<{
    access_token: string;
    refresh_token: string;
    user: { email: string; name: string; avatar_url?: string };
  }>;
}

export const authOptions: NextAuthOptions = {
  session: { strategy: "jwt" },
  providers: [
    GoogleProvider({
      clientId: process.env.GOOGLE_CLIENT_ID || "",
      clientSecret: process.env.GOOGLE_CLIENT_SECRET || "",
      allowDangerousEmailAccountLinking: true,
    }),
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        email: { label: "Email", type: "email" },
        password: { label: "Password", type: "password" },
      },
      async authorize(creds) {
        const email = creds?.email as string;
        const password = creds?.password as string;
        if (!email || !password) return null;
        const out = await loginWithBackend(email, password);
        if (!out) return null;
        return {
          id: email,
          email,
          name: out.user.name,
          image: out.user.avatar_url,
          backendAccess: out.access_token,
          backendRefresh: out.refresh_token,
        } as any;
      },
    }),
  ],
  callbacks: {
    async jwt({ token, user, account }) {
      if (account?.provider === "google" && (account as any).id_token) {
        const ex = await exchangeGoogleIdToken(
          (account as any).id_token as string
        );
        if (ex) {
          token.backendAccess = ex.access_token;
          token.backendRefresh = ex.refresh_token;
          token.picture = ex.user.avatar_url;
          token.name = ex.user.name;
          token.email = ex.user.email;
        }
      }
      if (user && (user as any).backendAccess) {
        token.backendAccess = (user as any).backendAccess;
        token.backendRefresh = (user as any).backendRefresh;
      }
      return token;
    },
    async session({ session, token }) {
      (session as any).backendAccess = (token as any).backendAccess;
      (session as any).backendRefresh = (token as any).backendRefresh;
      if (token.picture && session.user)
        session.user.image = token.picture as string;
      return session;
    },
  },
  pages: { signIn: "/login" },
};

export default NextAuth(authOptions);
