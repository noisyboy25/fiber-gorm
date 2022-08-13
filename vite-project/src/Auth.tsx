import React, { useState } from 'react';

const Auth = () => {
  enum AuthMode {
    Login = 'login',
    Register = 'register',
  }

  const [auth, setAuth] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

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

    setAuth(data.username);
  };

  return (
    <div>
      {auth || (
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
          <button>Login</button>
          <button type="button" onClick={(e) => submit(e, AuthMode.Register)}>
            Register
          </button>
        </form>
      )}
    </div>
  );
};

export default Auth;
