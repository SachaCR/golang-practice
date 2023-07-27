import React from 'react'
import ReactDOM from 'react-dom/client'
import { Auth0Provider } from '@auth0/auth0-react';

import App from './App.tsx'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <Auth0Provider
          domain="dev-golang-test.eu.auth0.com"
          clientId="KiigfCeTUPYc9SnQInAfN0RPkne2O3Wd"
          authorizationParams={{
            redirect_uri: window.location.origin
          }}
        >
      <App />
      </Auth0Provider>
  </React.StrictMode>,
)
