import Link from 'next/link'
import ThemeToggle from './theme-toggle'
import UserMenu from './user-menu'
export default function Navbar() {
  return (
    <header className="sticky top-0 z-40 w-full border-b border-border bg-bg/80 backdrop-blur-sm">
      <nav className="container flex h-14 items-center justify-between">
        <Link href="/" className="group inline-flex items-center gap-2">
          <span className="inline-block h-2 w-2 rounded-full bg-primary transition-transform group-hover:scale-125" />
          <span className="text-lg font-semibold tracking-tight">MagicStream</span>
        </Link>
        <div className="flex items-center gap-2"><ThemeToggle /><UserMenu /></div>
      </nav>
    </header>
  )
}
