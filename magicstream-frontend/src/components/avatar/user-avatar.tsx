'use client'
import Image from 'next/image'
import { stringToHsl, cn } from '@/lib/utils'
function initials(name?: string | null, email?: string | null) {
  const src = name || email || '?'
  const parts = src.trim().split(' ')
  if (parts.length >= 2) return (parts[0][0] + parts[1][0]).toUpperCase()
  return src.slice(0,2).toUpperCase()
}
export default function UserAvatar({ name, email, image, className }: { name?: string | null; email?: string | null; image?: string | null; className?: string }) {
  const bg = stringToHsl((email || name || 'user') as string)
  return (
    <div className={cn('inline-flex h-9 w-9 items-center justify-center overflow-hidden rounded-full ring-1 ring-border', className)} title={name || email || 'User'}>
      {image ? (
        <Image src={image} alt={name || 'avatar'} width={36} height={36} className="h-full w-full object-cover" />
      ) : (
        <div className="h-full w-full grid place-items-center text-xs font-semibold" style={{ background: `linear-gradient(135deg, ${bg}, color-mix(in oklab, ${bg}, white 20%))`, color: 'white' }}>{initials(name, email)}</div>
      )}
    </div>
  )
}
