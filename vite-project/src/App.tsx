import React, { useMemo, useState } from 'react';
import './App.css';
import Auth from './Auth';
import TodoList from './TodoList';

export const AuthContext = React.createContext({
  auth: '',
  setAuth: (_a: string) => {
    return;
  },
});

function App() {
  const [auth, setAuth] = useState('');
  const value = useMemo(
    () => ({ auth, setAuth: (a: string) => setAuth(a) }),
    [auth]
  );

  return (
    <div className="App">
      <AuthContext.Provider value={value}>
        <Auth />
        <TodoList />
      </AuthContext.Provider>
    </div>
  );
}

export default App;
