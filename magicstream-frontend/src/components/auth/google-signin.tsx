'use client'
import { signIn } from 'next-auth/react'
import { Button } from '@/components/ui/button'

export default function GoogleSignIn({text = 'Continue with Google'}:{text?:string}){
  return (
    <Button type="button" variant="outline" className="w-full gap-2" onClick={()=>signIn('google', { callbackUrl: '/' })}>
      <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 48 48" aria-hidden="true" focusable="false">
        <path fill="#FFC107" d="M43.611,20.083H42V20H24v8h11.303c-1.649,4.657-6.08,8-11.303,8c-6.627,0-12-5.373-12-12 s5.373-12,12-12c3.059,0,5.842,1.154,7.961,3.039l5.657-5.657C33.201,6.053,28.791,4,24,4C12.955,4,4,12.955,4,24 s8.955,20,20,20s20-8.955,20-20C44,22.659,43.862,21.35,43.611,20.083z"/>
        <path fill="#FF3D00" d="M6.306,14.691l6.571,4.819C14.655,15.108,18.961,12,24,12c3.059,0,5.842,1.154,7.961,3.039l5.657-5.657 C33.201,6.053,28.791,4,24,4C16.318,4,9.656,8.337,6.306,14.691z"/>
        <path fill="#4CAF50" d="M24,44c4.717,0,9.023-1.809,12.285-4.756l-5.671-4.773C28.543,35.341,26.396,36,24,36 c-5.202,0-9.616-3.317-11.279-7.946l-6.538,5.044C9.49,39.556,16.227,44,24,44z"/>
        <path fill="#1976D2" d="M43.611,20.083H42V20H24v8h11.303c-0.792,2.237-2.231,4.166-4.089,5.471 c0.001-0.001,0.002-0.001,0.003-0.002l5.671,4.773C35.549,38.789,44,32,44,24C44,22.659,43.862,21.35,43.611,20.083z"/>
      </svg>
      {text}
    </Button>
  )
}
