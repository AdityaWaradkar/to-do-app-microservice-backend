import React, { useState, useEffect } from "react";

const TodoComponent = () => {
  const [todos, setTodos] = useState([]);
  const [newTodo, setNewTodo] = useState({
    title: "",
    description: "",
    completed: false,
  });
  const [editTodo, setEditTodo] = useState(null);
  const [editText, setEditText] = useState("");
  const [userID, setUserID] = useState(localStorage.getItem("userID"));

  useEffect(() => {
    if (userID) {
      fetchTodos();
    }
  }, [userID]);

  const fetchTodos = async () => {
    try {
      const response = await fetch(
        `http://13.127.32.151:8082/api/todos/fetch?userID=${userID}`
      );
      const data = await response.json();

      if (Array.isArray(data)) {
        setTodos(data);
      } else {
        console.error("Error: Response data is not an array", data);
        setTodos([]);
      }
    } catch (error) {
      console.error("Error fetching todos:", error);
      setTodos([]);
    }
  };

  const handleAddTodo = async () => {
    if (!newTodo.title) return;

    const newTodoItem = {
      title: newTodo.title,
      description: newTodo.description,
      completed: newTodo.completed,
      user_id: userID,
    };

    try {
      const response = await fetch("http://13.127.32.151:8082/api/todos", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newTodoItem),
      });

      if (response.ok) {
        fetchTodos();
        setNewTodo({ title: "", description: "", completed: false });
      } else {
        console.error("Failed to add todo");
      }
    } catch (error) {
      console.error("Error adding todo:", error);
    }
  };

  const handleEditTodo = (id, title, description) => {
    setEditTodo(id);
    setEditText(title);
  };

  // Handle saving edited todo
  const handleSaveEdit = async () => {
    if (!editText) return;

    const updatedTodo = {
      id: editTodo,
      title: editText,
      description: newTodo.description,
      completed: false,
      user_id: userID,
    };

    try {
      const response = await fetch(
        `http://13.127.32.151:8082/api/todos/${editTodo}`,
        {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(updatedTodo),
        }
      );

      if (response.ok) {
        fetchTodos();
        setEditTodo(null);
        setEditText("");
      } else {
        console.error("Failed to save edit");
      }
    } catch (error) {
      console.error("Error updating todo:", error);
    }
  };

  const handleToggleComplete = async (id, currentStatus) => {
    try {
      const updatedTodo = {
        id: id,
        title: todos.find((todo) => todo.id === id).title,
        description: todos.find((todo) => todo.id === id).description,
        completed: !currentStatus,
        user_id: userID,
      };

      const response = await fetch(
        `http://13.127.32.151:8082/api/todos/${id}`,
        {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(updatedTodo),
        }
      );

      if (response.ok) {
        fetchTodos();
      } else {
        console.error("Failed to toggle completion status");
      }
    } catch (error) {
      console.error("Error toggling completion status:", error);
    }
  };

  const handleDeleteTodo = async (id) => {
    try {
      const response = await fetch(
        `http://13.127.32.151:8082/api/todos/${id}`,
        {
          method: "DELETE",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ user_id: userID }),
        }
      );

      if (response.ok) {
        fetchTodos();
      } else {
        console.error("Failed to delete todo");
      }
    } catch (error) {
      console.error("Error deleting todo:", error);
    }
  };

  return (
    <div className="max-w-4xl mx-auto p-6 bg-white rounded-lg shadow-lg">
      <h2 className="text-3xl font-semibold mb-6 text-center">
        Your To-Do List
      </h2>

      <div className="mb-6 flex flex-col md:flex-row items-center justify-between space-y-4 md:space-y-0">
        <input
          type="text"
          value={newTodo.title}
          onChange={(e) => setNewTodo({ ...newTodo, title: e.target.value })}
          className="border border-gray-300 p-3 w-full md:w-3/4 rounded-md"
          placeholder="Add a new to-do"
        />
        <button
          onClick={handleAddTodo}
          className="bg-blue-500 text-white px-6 py-3 rounded-md hover:bg-blue-600 mt-4 md:mt-0"
        >
          Add
        </button>
      </div>

      {Array.isArray(todos) && todos.length > 0 ? (
        <ul className="space-y-4">
          {todos.map((todo) => (
            <li
              key={todo.id}
              className="flex justify-between items-center p-4 bg-gray-50 rounded-lg shadow-sm"
            >
              <div className="flex items-center space-x-4">
                <input
                  type="checkbox"
                  checked={todo.completed}
                  onChange={() => handleToggleComplete(todo.id, todo.completed)}
                  className="mr-4"
                />
                {editTodo === todo.id ? (
                  <input
                    type="text"
                    value={editText}
                    onChange={(e) => setEditText(e.target.value)}
                    className="border border-gray-300 p-3 w-64 rounded-md"
                  />
                ) : (
                  <span
                    className={`text-lg ${
                      todo.completed ? "line-through text-gray-500" : ""
                    }`}
                  >
                    {todo.title}
                  </span>
                )}
              </div>

              <div className="space-x-3 flex items-center">
                {editTodo === todo.id ? (
                  <button
                    onClick={handleSaveEdit}
                    className="bg-green-500 text-white px-4 py-2 rounded-md hover:bg-green-600"
                  >
                    Save
                  </button>
                ) : (
                  <button
                    onClick={() =>
                      handleEditTodo(todo.id, todo.title, todo.description)
                    }
                    className="bg-yellow-500 text-white px-4 py-2 rounded-md hover:bg-yellow-600"
                  >
                    Edit
                  </button>
                )}

                <button
                  onClick={() => handleDeleteTodo(todo.id)}
                  className="bg-red-500 text-white px-4 py-2 rounded-md hover:bg-red-600"
                >
                  Delete
                </button>
              </div>
            </li>
          ))}
        </ul>
      ) : (
        <p className="text-center text-gray-600">
          No tasks yet. Start by adding some!
        </p>
      )}
    </div>
  );
};

export default TodoComponent;
