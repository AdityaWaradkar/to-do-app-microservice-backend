import React from "react";
import { useNavigate } from "react-router-dom";

const Navbar = () => {
  const navigate = useNavigate();

  // Ensure that you're processing the response as an object, not an array
  const handleLogout = () => {
    // Clear token and userID from localStorage
    localStorage.removeItem("userID");
    localStorage.removeItem("token");
  
    // Redirect user to the login page
    navigate("/");
  };
  

  return (
    <nav className="bg-blue-500 text-white p-4 flex justify-between items-center">
      <div className="text-lg font-bold">To-Do List</div>
      <div>
        <button
          onClick={handleLogout}
          className="bg-red-500 text-white p-2 rounded hover:bg-red-600"
        >
          Logout
        </button>
      </div>
    </nav>
  );
};

export default Navbar;
