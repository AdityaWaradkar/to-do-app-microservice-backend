import React from "react";
import Navbar from "./Navbar"; // Import the Navbar component
import TodoComponent from "./Todo";

const Home = () => {
  return (
    <div className="min-h-screen bg-gray-100">
      {/* Full height screen and background color */}
      <Navbar /> {/* Include the Navbar in the Home page */}
      <div className="flex flex-col items-center justify-center p-6 sm:p-12 md:p-16">
        <div className="w-full max-w-5xl">
          {/* Make sure TodoComponent doesn't stretch too wide */}
          <TodoComponent />
        </div>
      </div>
    </div>
  );
};

export default Home;
