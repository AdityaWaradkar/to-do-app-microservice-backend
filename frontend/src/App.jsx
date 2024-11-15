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
        <Route path="/login" element={<Navigate to="/" />} />{" "}
        {/* Redirect from /login to root */}
      </Routes>
    </Router>
  );
};

export default App;