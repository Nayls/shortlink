import React from 'react'
import { useSelector } from 'react-redux'
import Router from 'next/router'

export default function withAuthSync(Child: any) {
  return (props?: any) => {
    // checks whether we are on client / browser or server.
    if (typeof window !== 'undefined') {
      // @ts-ignore
      const session = useSelector((state) => state.session)

      if (!session.kratos.active) {
        Router.push('/auth/login')
        return null
      }
    }

    // If this is an token we just render the component that was passed with all its props
    return <Child {...props} />
  }
}
