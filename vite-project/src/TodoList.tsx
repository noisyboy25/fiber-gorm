import React, { useEffect, useState } from 'react';
import './Todo.css';

type Todo = {
  id: number;
  text: string;
};

const TodoList = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [text, setText] = useState('');

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
    await fetch('/api/todos', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ text: inputText }),
    });
    await getTodos();
  };
  const deleteTodos = async (i: number) => {
    const res = await fetch(`/api/todos/${i}`, {
      method: 'DELETE',
    });
    if (res.ok) getTodos();
  };

  useEffect(() => {
    getTodos();
  }, []);

  return (
    <div>
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
      <ul>
        {todos.map((todo) => (
          <React.Fragment key={todo.id}>
            <p>{todo.text}</p>
            <button onClick={() => deleteTodos(todo.id)}>x</button>
          </React.Fragment>
        ))}
      </ul>
    </div>
  );
};

export default TodoList;
