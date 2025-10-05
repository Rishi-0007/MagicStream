'use client'
import { useState } from 'react'
import { useSession, signIn, signOut } from 'next-auth/react'
import Link from 'next/link'
import UserAvatar from '@/components/avatar/user-avatar'
import { Button } from '@/components/ui/button'
export default function UserMenu() {
  const { data: session, status } = useSession()
  const [open, setOpen] = useState(false)
  if (status === 'loading') { return <div className="h-9 w-20 rounded-full bg-elevated animate-pulse" /> }
  if (!session) { return (<div className="flex items-center gap-2"><Button variant="outline" onClick={() => signIn(undefined, { callbackUrl: '/' })}>Login</Button><Link href="/register" className="text-sm text-muted hover:text-text">Register</Link></div>) }
  const name = session.user?.name; const email = session.user?.email; const image = session.user?.image
  return (<div className="relative">
    <button onClick={() => setOpen((o) => !o)} className="rounded-full focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring">
      <UserAvatar name={name} email={email ?? undefined} image={image ?? undefined} />
    </button>
    {open && (<div className="absolute right-0 mt-2 w-56 overflow-hidden rounded-[var(--radius)] border border-border bg-surface shadow-xl">
      <div className="px-3 py-2 text-sm"><div className="font-medium truncate">{name || email}</div><div className="text-muted truncate">{email}</div></div>
      <div className="border-t border-border" />
      <ul className="p-1 text-sm">
        <li><Link onClick={()=>setOpen(false)} className="block rounded px-3 py-2 hover:bg-elevated" href="/account">My Account</Link></li>
        <li><button className="w-full text-left rounded px-3 py-2 hover:bg-elevated" onClick={() => signOut({ callbackUrl: '/' })}>Sign out</button></li>
      </ul>
    </div>)}
  </div>)
}
