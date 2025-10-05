'use client'
import { useTheme } from 'next-themes'
import { Moon, Sun } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { useEffect, useState } from 'react'
export default function ThemeToggle() {
  const { theme, setTheme, systemTheme } = useTheme()
  const [mounted, setMounted] = useState(false)
  useEffect(() => setMounted(true), [])
  const current = theme === 'system' ? systemTheme : theme
  if (!mounted) return null
  return <Button variant="ghost" size="icon" aria-label="Toggle theme" onClick={() => setTheme(current === 'dark' ? 'light' : 'dark')} title="Toggle theme">{current === 'dark' ? <Sun className="h-5 w-5" /> : <Moon className="h-5 w-5" />}</Button>
}
