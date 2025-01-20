import { useState } from "react";
import { useNavigate } from "react-router-dom";

const Register = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [email, setEmail] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const handleRegister = async (e) => {
    e.preventDefault();
    if (!email || !password || !username) {
      setError("All fields are required");
      return;
    }
    setLoading(true);
    setError(null);
    const userData = { username, password, email };

    try {
      const response = await fetch(
        "http://13.127.32.151:8081/api/user/register",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(userData),
        }
      );

      const contentType = response.headers.get("content-type");
      const textResponse = await response.text();

      if (response.ok) {
        if (contentType && contentType.includes("application/json")) {
          const jsonResponse = JSON.parse(textResponse);
          navigate("/");
        } else {
          setError("Unexpected response format.");
        }
      } else {
        const errorData =
          contentType && contentType.includes("application/json")
            ? JSON.parse(textResponse)
            : { error: textResponse };
        setError(errorData.error || "Registration failed.");
      }
    } catch (error) {
      console.error("Error during registration:", error);
      setError("An unexpected error occurred.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <form
        onSubmit={handleRegister}
        className="bg-white p-6 rounded shadow-md w-80"
      >
        <h2 className="text-lg font-bold mb-4">Register</h2>
        {error && <p className="text-red-500 mb-4">{error}</p>}{" "}
        {/* Error message */}
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
          className={`bg-blue-500 text-white p-2 w-full rounded hover:bg-blue-600 ${
            loading && "opacity-50 cursor-not-allowed"
          }`}
          disabled={loading}
        >
          {loading ? "Registering..." : "Register"}
        </button>
        <p className="mt-4 text-gray-600">
          Already have an account?{" "}
          <a href="/" className="text-blue-500 hover:text-blue-700">
            Login here
          </a>
        </p>
      </form>
    </div>
  );
};

export default Register;
