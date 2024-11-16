import React from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";
import Login from "./components/Login";
import Register from "./components/Register";
import Home from "./components/Home";

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />{" "}
        {/* Display Login on root path */}
        <Route path="/home" element={<Home />} /> {/* Home page at /home */}
        <Route path="/register" element={<Register />} />
        {/* If logged in, redirect to /home, else show /login */}
        {/* You can add a simple check for authentication here */}
        <Route path="/login" element={<Navigate to="/" />} />
      </Routes>
    </Router>
  );
};

export default App;
