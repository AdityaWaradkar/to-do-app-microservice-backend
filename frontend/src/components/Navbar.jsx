import React from "react";
import { useNavigate } from "react-router-dom";

const Navbar = () => {
  const navigate = useNavigate();

  const handleLogout = async () => {
    try {
      // Make a request to log the user out (clear session in backend)
      const response = await fetch(
        "http://13.203.79.155:8081/api/user/logout",
        {
          method: "POST",
          credentials: "include", // Ensure cookies are sent with the request
        }
      );

      if (response.ok) {
        // Redirect user to the login page after successful logout
        navigate("/login");
      } else {
        console.error("Logout failed");
      }
    } catch (error) {
      console.error("Error logging out:", error);
    }
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
