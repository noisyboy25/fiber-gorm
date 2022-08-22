import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './App';

const Auth = () => {
  enum AuthMode {
    Login = 'logIn',
    Register = 'register',
  }

  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const { auth, setAuth } = useContext(AuthContext);

  const logOut = () => {
    localStorage.removeItem('auth');
    setAuth('');
  };

  useEffect(() => {
    (() => {
      setAuth(localStorage.getItem('auth') ?? '');
    })();
  }, []);

  const submit = async (event: React.FormEvent, mode: AuthMode) => {
    event.preventDefault();
    const res = await fetch(`/auth/${mode}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password }),
    });
    if (!res.ok) {
      return console.log(await res.text());
    }
    const data = await res.json();
    console.log(data);

    localStorage.setItem('auth', data.auth);
    setAuth(localStorage.getItem('auth') ?? '');
  };

  return (
    <div>
      {!auth ? (
        <form onSubmit={(e) => submit(e, AuthMode.Login)}>
          <div>
            <input
              type="text"
              name="username"
              placeholder="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>
          <div>
            <input
              type="password"
              name="password"
              placeholder="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <button>Log In</button>
          <button type="button" onClick={(e) => submit(e, AuthMode.Register)}>
            Register
          </button>
        </form>
      ) : (
        <div>
          Logged in as{' '}
          <span style={{ fontWeight: 600 }}>{auth.split(':')[0]}</span>
          <button onClick={logOut}>Log Out</button>
        </div>
      )}
    </div>
  );
};

export default Auth;
