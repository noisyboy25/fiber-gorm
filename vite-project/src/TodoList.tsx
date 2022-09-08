import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './App';
import './Todo.css';

type Todo = {
  id: number;
  text: string;
  userId: string;
  user: any;
};

const TodoList = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [text, setText] = useState('');

  const { auth } = useContext(AuthContext);

  const clear = () => {
    setText('');
  };

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    await addTodos(text);
    clear();
  };

  const getTodos = async () => {
    const res = await fetch('/api/todos');
    const json = await res.json();
    setTodos(json.todos);
  };

  const addTodos = async (inputText: string) => {
    const res = await fetch('/api/todos', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Basic ${localStorage.getItem('auth')}`,
      },
      body: JSON.stringify({ text: inputText }),
    });
    if (!res.ok) {
      console.log(await res.text());
    }
    await getTodos();
  };

  const deleteTodos = async (i: number) => {
    const res = await fetch(`/api/todos/${i}`, {
      method: 'DELETE',
      headers: {
        Authorization: `Basic ${localStorage.getItem('auth')}`,
      },
    });
    if (res.ok) getTodos();
    else console.log(await res.text());
  };

  useEffect(() => {
    getTodos();
  }, []);

  return (
    <div>
      {auth && (
        <form onSubmit={onSubmit}>
          <input
            type="text"
            name="text"
            id="text"
            onChange={(event) => setText(event.target.value)}
            value={text}
            autoComplete="off"
          />
        </form>
      )}
      <div className="list">
        {todos.map((todo) => (
          <React.Fragment key={todo.id}>
            <p className="author">{todo.user?.username || 'unknown'}</p>
            <p>{todo.text}</p>
            {auth && <button onClick={() => deleteTodos(todo.id)}>x</button>}
          </React.Fragment>
        ))}
      </div>
    </div>
  );
};

export default TodoList;
