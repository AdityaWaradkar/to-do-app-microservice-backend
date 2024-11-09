import React, { useState, useEffect } from "react";

const TodoComponent = () => {
  const [todos, setTodos] = useState([]); // Array of to-dos
  const [newTodo, setNewTodo] = useState(""); // To store input for new todo
  const [editTodo, setEditTodo] = useState(null); // To track which todo is being edited
  const [editText, setEditText] = useState(""); // To store edited text

  // Fetch all todos when the component mounts
  useEffect(() => {
    fetchTodos();
  }, []);

  // Fetch todos from backend API
  const fetchTodos = async () => {
    try {
      const response = await fetch("http://localhost:8082/todos");
      const data = await response.json();
      setTodos(data);
    } catch (error) {
      console.error("Error fetching todos:", error);
    }
  };

  // Handle adding a new todo
  const handleAddTodo = async () => {
    if (!newTodo) return; // Do nothing if input is empty

    const newTodoItem = {
      text: newTodo,
      completed: false,
    };

    try {
      const response = await fetch("http://localhost:8082/todos", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newTodoItem),
      });

      if (response.ok) {
        fetchTodos(); // Refresh the list of todos after adding
        setNewTodo(""); // Clear input after adding
      } else {
        console.error("Failed to add todo");
      }
    } catch (error) {
      console.error("Error adding todo:", error);
    }
  };

  // Handle editing a todo
  const handleEditTodo = (id, text) => {
    setEditTodo(id);
    setEditText(text); // Set current todo text to edit
  };

  // Handle saving edited todo
  const handleSaveEdit = async () => {
    if (!editText) return; // Do nothing if input is empty

    const updatedTodo = {
      text: editText,
      completed: false, // You can change this if editing completion status
    };

    try {
      const response = await fetch(`http://localhost:8082/todos/${editTodo}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(updatedTodo),
      });

      if (response.ok) {
        fetchTodos(); // Refresh the list after editing
        setEditTodo(null); // Reset edit state
        setEditText(""); // Clear edit text
      } else {
        console.error("Failed to save edit");
      }
    } catch (error) {
      console.error("Error updating todo:", error);
    }
  };

  // Handle toggling the completion status of a todo
  const handleToggleComplete = async (id, currentStatus) => {
    try {
      const updatedTodo = {
        text: todos.find((todo) => todo.id === id).text, // Keep the text the same
        completed: !currentStatus, // Toggle completion status
      };

      const response = await fetch(`http://localhost:8082/todos/${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(updatedTodo),
      });

      if (response.ok) {
        fetchTodos(); // Refresh the list after toggling
      } else {
        console.error("Failed to toggle completion status");
      }
    } catch (error) {
      console.error("Error toggling completion status:", error);
    }
  };

  // Handle deleting a todo
  const handleDeleteTodo = async (id) => {
    try {
      const response = await fetch(`http://localhost:8082/todos/${id}`, {
        method: "DELETE",
      });

      if (response.ok) {
        fetchTodos(); // Refresh the list after deletion
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

      {/* Add new To-Do */}
      <div className="mb-6 flex flex-col md:flex-row items-center justify-between space-y-4 md:space-y-0">
        <input
          type="text"
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
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

      {/* Todo List */}
      {todos.length === 0 ? (
        <p className="text-center text-gray-600">
          No tasks yet. Start by adding some!
        </p>
      ) : (
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
                    {todo.text}
                  </span>
                )}
              </div>

              {/* Action Buttons */}
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
                    onClick={() => handleEditTodo(todo.id, todo.text)}
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
      )}
    </div>
  );
};

export default TodoComponent;
