'use client'
import GoogleSignIn from '@/components/auth/google-signin'
import CredentialsForm from '@/components/auth/credentials-form'
import Link from 'next/link'
export default function LoginPage(){
  return (<div className="mx-auto w-full max-w-sm space-y-4">
    <div className="space-y-1 text-center"><h2>Welcome back</h2><p className="text-muted text-sm">Sign in to continue</p></div>
    <GoogleSignIn />
    <div className="relative py-2"><div className="absolute inset-0 flex items-center"><span className="w-full border-t border-border" /></div><div className="relative flex justify-center"><span className="bg-bg px-2 text-xs text-muted">or</span></div></div>
    <CredentialsForm />
    <p className="text-center text-sm text-muted">No account? <Link href="/register" className="text-primary">Register</Link></p>
  </div>)
}
