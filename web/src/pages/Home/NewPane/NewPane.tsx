import { PaneHeading, PaneSubheading } from '../../../components/Panes/Panes'

export default function NewPane() {
  return (
    <div>
      <PaneHeading>Create a new notebook</PaneHeading>
      <PaneSubheading>
        Those settings <strong>can not</strong> be changed later.
      </PaneSubheading>
    </div>
  )
}
