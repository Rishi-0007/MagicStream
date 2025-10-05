import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'

export default function HomePage() {
  return (
    <section className="space-y-8">
      <div className="space-y-2">
        <h1>Stream smarter.</h1>
        <p className="text-muted">Beautiful, fast, and accessible. Your cinematic hub — coming next.</p>
      </div>

      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
        {[1,2,3].map((i) => (
          <Card key={i}>
            <CardHeader>
              <CardTitle>Feature {i}</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-muted">This is a placeholder card to validate styling and tokens.</p>
              <div className="mt-4 flex gap-2">
                <Button>Primary</Button>
                <Button variant="outline">Outline</Button>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>
    </section>
  )
}
