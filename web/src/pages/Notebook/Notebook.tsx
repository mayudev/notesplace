import { SerializedError } from '@reduxjs/toolkit'
import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { useAppDispatch, useAppSelector } from '../../app/hooks'
import {
  fetchNotebook,
  selectNotebookData,
} from '../../features/notebook/notebook.slice'
import Layout from './Layout/Layout'
import NotebookError from './NotebookError/NotebookError'

type Params = {
  id: string
}

export default function Notebook() {
  let params = useParams<Params>()
  const dispatch = useAppDispatch()

  const notebook = useAppSelector(selectNotebookData)
  const status = useAppSelector(state => state.notebook.status)

  const [errorMessage, setErrorMessage] = useState('')

  useEffect(() => {
    async function fetchData(id: string) {
      try {
        await dispatch(
          fetchNotebook({
            id,
            jwt: '', // TODO
          })
        ).unwrap()
      } catch (e) {
        const err = e as SerializedError
        switch (err.code) {
          case '401':
            setErrorMessage('Not authorized.')
            // show password prompt
            break
          case '404':
            setErrorMessage('Notebook not found.')
            break
          case '500':
            setErrorMessage('Internal server error.')
            break
          default:
            setErrorMessage('An unknown error occurred.')
            break
        }
      }
    }

    // If current notebook isn't already present in state, fetch it.
    if (params.id !== notebook.id) {
      fetchData(params.id!)
    }
  }, [dispatch, params.id, notebook.id])

  const display = () => {
    switch (status) {
      case 'failed':
        return <NotebookError>{errorMessage}</NotebookError>
      case 'pending':
        return <div>loading</div>
      case 'idle':
      case 'succeeded':
        return <div>name: {notebook.name}</div>
    }
  }

  return (
    <>
      <Layout />
      {display()}
    </>
  )
}
