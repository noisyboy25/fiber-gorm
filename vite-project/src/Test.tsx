import React, { useEffect, useState } from 'react';

const Test = () => {
  const [list, setList] = useState<string[]>([]);
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

  const addList = async (text: string) => {
    await fetch('/api/list', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ text: text }),
    });
    await getList();
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
        />
      </form>
      {list.map((el) => (
        <div>{el}</div>
      ))}
    </div>
  );
};

export default Test;
