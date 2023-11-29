import CssBaseline from '@mui/material/CssBaseline';
import { GoogleOAuthProvider } from '@react-oauth/google';
import React from 'react';
import ReactDOM from 'react-dom/client';
import {
  createBrowserRouter,
  RouterProvider
} from "react-router-dom";

import App from './App';
import Login from './Login';

const router = createBrowserRouter([
  {
    path: "/app",
    element: <App />
  },
  {
    path: "/",
    element: <Login />
  },
]);

const root = ReactDOM.createRoot(document.getElementById('root'));

root.render(
  <React.StrictMode>
    <GoogleOAuthProvider clientId="34507644075-mil8kn8aubbhsrkuc234probjer4uigl.apps.googleusercontent.com">
      <CssBaseline />
      <RouterProvider router={router} />
    </GoogleOAuthProvider>
  </React.StrictMode>
);



