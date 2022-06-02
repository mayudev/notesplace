import { SerializedError } from '@reduxjs/toolkit'
import { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { useAppDispatch, useAppSelector } from '../../app/hooks'
import PasswordPrompt from '../../components/PasswordPrompt/PasswordPrompt'
import Spinner from '../../components/Spinner/Spinner'
import { clearToken } from '../../features/global/global.slice'
import {
  clearNotebook,
  fetchNotebook,
  selectNotebookData,
} from '../../features/notebook/notebook.slice'
import Layout from './Layout/Layout'
import { Center } from './Notebook.styles'
import NotebookError from './NotebookError/NotebookError'
import Notes from './Notes/Notes'

type Params = {
  id: string
}

export default function Notebook() {
  let params = useParams<Params>()
  const navigate = useNavigate()
  const dispatch = useAppDispatch()

  const notebook = useAppSelector(selectNotebookData)
  const status = useAppSelector(state => state.notebook.status)

  const [errorMessage, setErrorMessage] = useState('')
  const [showPasswordPrompt, setShowPasswordPrompt] = useState(false)

  const fetchData = async (id: string, token: string) => {
    try {
      await dispatch(
        fetchNotebook({
          id,
          jwt: token, // TODO
        })
      ).unwrap()
    } catch (e) {
      const err = e as SerializedError
      switch (err.code) {
        case '401':
          setErrorMessage('Not authorized')
          setShowPasswordPrompt(true)
          break
        case '404':
          setErrorMessage('Notebook not found')
          break
        case '500':
          setErrorMessage('Internal server error')
          break
        default:
          setErrorMessage('An unknown error occurred')
          break
      }
    }
  }
  useEffect(() => {
    // If current notebook isn't already present in state, fetch it.
    if (params.id !== notebook.id) {
      fetchData(params.id!, '')
    }
  }, [params.id, notebook.id])

  const passwordEntered = async (success: boolean, token: string) => {
    setShowPasswordPrompt(false)

    if (!success) return

    fetchData(params.id!, token)
  }

  const logout = () => {
    // CLear the token
    dispatch(clearToken())
    dispatch(clearNotebook())

    // Redirect to home page
    navigate('/')
  }

  const display = () => {
    switch (status) {
      case 'failed':
        if (showPasswordPrompt) {
          return (
            <PasswordPrompt notebook={params.id!} onSubmit={passwordEntered} />
          )
        }
        return <NotebookError>{errorMessage}</NotebookError>
      case 'idle':
      case 'pending':
        return (
          <Center>
            <Spinner width={50} borderWidth={5} primary />
          </Center>
        )
      case 'succeeded':
        return <Notes />
    }
  }

  return (
    <>
      <Layout onClose={logout} />
      {display()}
    </>
  )
}
