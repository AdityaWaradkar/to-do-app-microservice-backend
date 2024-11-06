import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const Register = () => {
  const [username, setUsername] = useState(""); // For username
  const [password, setPassword] = useState(""); // For password
  const [email, setEmail] = useState(""); // For email
  const navigate = useNavigate();

  const handleRegister = async (e) => {
    e.preventDefault();
    const userData = {
      username,
      password,
      email,
    };

    try {
      const response = await fetch("http://localhost:8081/api/users/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userData),
      });

      const contentType = response.headers.get("content-type");
      const textResponse = await response.text(); // Get the text response

      console.log("Response:", response);
      console.log("Response Text:", textResponse);

      if (response.ok) {
        if (contentType && contentType.includes("application/json")) {
          const jsonResponse = JSON.parse(textResponse);
          console.log("Registered User:", jsonResponse);
        }
        navigate("/login");
      } else {
        const errorData =
          contentType && contentType.includes("application/json")
            ? JSON.parse(textResponse)
            : { error: textResponse };
        alert(errorData.error || "Registration failed.");
      }
    } catch (error) {
      console.error("Error during registration:", error);
      alert("An unexpected error occurred.");
    }
  };

  return (
    <div className="flex items-center justify-center h-screen bg-gray-100">
      <form
        onSubmit={handleRegister}
        className="bg-white p-6 rounded shadow-md w-80"
      >
        <h2 className="text-lg font-bold mb-4">Register</h2>
        <div className="mb-4">
          <label className="block text-gray-700">Username</label>
          <input
            type="text"
            className="border border-gray-300 p-2 w-full rounded"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Email</label>
          <input
            type="email"
            className="border border-gray-300 p-2 w-full rounded"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Password</label>
          <input
            type="password"
            className="border border-gray-300 p-2 w-full rounded"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button
          type="submit"
          className="bg-blue-500 text-white p-2 w-full rounded hover:bg-blue-600"
        >
          Register
        </button>
        <p className="mt-4 text-gray-600">
          Already have an account?{" "}
          <a href="/login" className="text-blue-500">
            Login
          </a>
        </p>
      </form>
    </div>
  );
};

export default Register;
