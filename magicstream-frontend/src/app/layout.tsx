import type { Metadata } from 'next'
import { Inter, Plus_Jakarta_Sans } from 'next/font/google'
import './globals.css'
import QueryProvider from '@/components/providers/query-provider'
import { ThemeProvider } from '@/components/providers/theme-provider'
import NextAuthProvider from '@/components/providers/session-provider'
import Navbar from '@/components/layout/navbar'
import Footer from '@/components/layout/footer'

const inter = Inter({ subsets: ['latin'], display: 'swap', variable: '--font-inter' })
const jakarta = Plus_Jakarta_Sans({ subsets: ['latin'], display: 'swap', variable: '--font-jakarta' })

export const metadata: Metadata = { title: 'MagicStream', description: 'A cinematic streaming experience.' }

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" suppressHydrationWarning className={`${inter.variable} ${jakarta.variable}`}>
      <body className="min-h-dvh bg-bg text-text antialiased">
        <ThemeProvider>
          <NextAuthProvider>
            <QueryProvider>
              <Navbar />
              <main className="container py-8">{children}</main>
              <Footer />
            </QueryProvider>
          </NextAuthProvider>
        </ThemeProvider>
      </body>
    </html>
  )
}
