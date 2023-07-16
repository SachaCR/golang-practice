import { useState } from 'react'
import { Auth0Provider, useAuth0 } from '@auth0/auth0-react';

import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'


  const LoginButton = () => {
    const { loginWithRedirect } = useAuth0();

    return <button onClick={() => loginWithRedirect()}>Log In</button>;
  };

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
      <Auth0Provider
          domain="dev-golang-test.eu.auth0.com"
          clientId="KiigfCeTUPYc9SnQInAfN0RPkne2O3Wd"
          authorizationParams={{
            redirect_uri: window.location.origin
          }}
        >
      <LoginButton></LoginButton>
      </Auth0Provider>
    </>
  )
}

export default App
