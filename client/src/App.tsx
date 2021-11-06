import React, { useEffect, useState } from "react";
import "./App.css";

type User = {
  username: string;
  password: string;
  created_on: string;
};

function isUser(param: any): param is User {
  return (
    param !== undefined &&
    param !== null &&
    typeof param === "object" &&
    typeof param.username === "string" &&
    typeof param.password === "string" &&
    typeof param.created_on === "string"
  );
}
function App() {
  const [credentials, setCredentials] = useState({
    username: "",
    password: "",
    created_on: "",
  });

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    try {
      const response = await fetch(`http://localhost:8080/register`, {
        method: "POST", // *GET, POST, PUT, DELETE, etc.
        body: JSON.stringify({
          username: credentials.username,
          password: credentials.password,
        }),
        credentials: "include",
      });

      // const response = await fetch("http://localhost:8080/logged-in", {
      //   method: "GET",
      //   credentials: "include",
      // });
      const data = await response.json();
      console.log(data);

      // if (isUser(data)) setCredentials(data);
    } catch (error: unknown) {
      console.error(error);
    }
  }

  function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
    e.preventDefault();
    setCredentials((prevState) => ({
      ...prevState,
      [e.target.name]: e.target.value,
    }));
  }

  useEffect(() => {
    console.log(credentials.created_on);
  }, [credentials.created_on]);

  return (
    <div className="App">
      <div>
        {credentials.username} - {credentials.password} -{" "}
      </div>
      <form onSubmit={handleSubmit}>
        <label>
          Username:
          <input
            type="text"
            name="username"
            onChange={handleChange}
            value={credentials.username}
          />
        </label>
        <label>
          Password:
          <input
            type="password"
            name="password"
            onChange={handleChange}
            value={credentials.password}
          />
        </label>
        <button type="submit" value="Submit">
          Submit
        </button>
      </form>
    </div>
  );
}

export default App;
