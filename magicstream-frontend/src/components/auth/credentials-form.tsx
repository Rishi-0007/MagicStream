'use client'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { signIn } from 'next-auth/react'

const schema = z.object({
  email: z.string().email(),
  password: z.string().min(6),
})

export default function CredentialsForm() {
  const { register, handleSubmit, formState: { errors, isSubmitting } } = useForm<z.infer<typeof schema>>({ resolver: zodResolver(schema) })
  const [error, setError] = useState<string|undefined>()

  const onSubmit = async (data: z.infer<typeof schema>) => {
    setError(undefined)
    const res = await signIn('credentials', { redirect: false, email: data.email, password: data.password })
    if (res?.error) setError('Invalid email or password')
    else window.location.href = '/'
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-3">
      <div>
        <Input placeholder="you@example.com" type="email" {...register('email')} />
        {errors.email && <p className="text-xs text-danger mt-1">{errors.email.message}</p>}
      </div>
      <div>
        <Input placeholder="••••••••" type="password" {...register('password')} />
        {errors.password && <p className="text-xs text-danger mt-1">{errors.password.message}</p>}
      </div>
      {error && <p className="text-sm text-danger">{error}</p>}
      <Button className="w-full" disabled={isSubmitting} type="submit">{isSubmitting ? 'Signing in...' : 'Sign in'}</Button>
    </form>
  )
}
