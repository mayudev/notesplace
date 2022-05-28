import { useParams } from 'react-router-dom'

type Params = {
  id: string
}

export default function Notebook() {
  let params = useParams<Params>()

  return (
    <div>
      <h1>Notebook {params.id!}!</h1>
    </div>
  )
}
