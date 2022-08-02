import React, { useEffect, useState } from 'react';
import './Test.css';

type Text = {
  id: number;
  text: string;
};

const Test = () => {
  const [list, setList] = useState<Text[]>([]);
  const [text, setText] = useState('');

  const clear = () => {
    setText('');
  };

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    await addList(text);
    clear();
  };

  const getList = async () => {
    const res = await fetch('/api/list');
    const json = await res.json();
    setList(json.list);
  };

  const addList = async (inputText: string) => {
    await fetch('/api/list', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ text: inputText }),
    });
    await getList();
  };
  const deleteList = async (i: number) => {
    const res = await fetch(`/api/list/${i}`, {
      method: 'DELETE',
    });
    if (res.ok) getList();
  };

  useEffect(() => {
    getList();
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
        {list.map((el) => (
          <React.Fragment key={el.id}>
            <p>{el.text}</p>
            <button onClick={() => deleteList(el.id)}>x</button>
          </React.Fragment>
        ))}
      </ul>
    </div>
  );
};

export default Test;
