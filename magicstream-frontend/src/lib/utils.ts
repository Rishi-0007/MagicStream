import { type ClassValue } from 'clsx'
import clsx from 'clsx'
import { twMerge } from 'tailwind-merge'
export function cn(...inputs: ClassValue[]) { return twMerge(clsx(inputs)) }
export function stringToHsl(str: string){ let h=0; for(let i=0;i<str.length;i++){ h = (h*31 + str.charCodeAt(i)) % 360 } return `hsl(${h} 70% 45%)` }
