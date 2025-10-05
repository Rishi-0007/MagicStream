'use client'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import api from '@/lib/axios'
import { signIn } from 'next-auth/react'
import Link from 'next/link'

const schema = z.object({ name: z.string().min(2), email: z.string().email(), password: z.string().min(6) })

export default function RegisterPage(){
  const { register, handleSubmit, formState: { errors, isSubmitting } } = useForm<z.infer<typeof schema>>({ resolver: zodResolver(schema) })
  const onSubmit = async (data: z.infer<typeof schema>) => {
    await api.post('/api/auth/register', data)
    const res = await signIn('credentials', { redirect: false, email: data.email, password: data.password })
    if (!res?.error) window.location.href = '/'
  }
  return (<div className="mx-auto w-full max-w-sm space-y-4">
    <div className="space-y-1 text-center"><h2>Create account</h2><p className="text-muted text-sm">It’s quick and easy.</p></div>
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-3">
      <div><Input placeholder="Your name" {...register('name')} />{errors.name && <p className="text-xs text-danger mt-1">{errors.name.message}</p>}</div>
      <div><Input placeholder="you@example.com" type="email" {...register('email')} />{errors.email && <p className="text-xs text-danger mt-1">{errors.email.message}</p>}</div>
      <div><Input placeholder="••••••••" type="password" {...register('password')} />{errors.password && <p className="text-xs text-danger mt-1">{errors.password.message}</p>}</div>
      <Button className="w-full" disabled={isSubmitting} type="submit">{isSubmitting ? 'Creating...' : 'Create account'}</Button>
    </form>
    <p className="text-center text-sm text-muted">Already have an account? <Link href="/login" className="text-primary">Login</Link></p>
  </div>)
}
