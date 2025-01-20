import React from "react";
import Navbar from "./Navbar";
import TodoComponent from "./Todo";

const Home = () => {
  return (
    <div className="min-h-screen bg-gray-100">
      <Navbar />
      <div className="flex flex-col items-center justify-center p-6 sm:p-12 md:p-16">
        <div className="w-full max-w-5xl">
          <TodoComponent />
        </div>
      </div>
    </div>
  );
};

export default Home;
