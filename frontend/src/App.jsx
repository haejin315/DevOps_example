import { useState } from "react";

function App() {
  const [a, setA] = useState("");
  const [b, setB] = useState("");
  const [result, setResult] = useState(null);
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    const formData = new URLSearchParams();
    formData.append("a", a);
    formData.append("b", b);

    try {
      const res = await fetch("http://localhost:8080/api/add", {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: formData,
      });

      if (!res.ok) {
        const msg = await res.text();
        throw new Error(msg);
      }

      const data = await res.json();
      setResult(data);
      setError("");
    } catch (err) {
      setResult(null);
      setError(err.message);
    }
  };

  return (
    <div style={{ padding: 20 }}>
      <h2>Add Two Numbers</h2>
      <form onSubmit={handleSubmit}>
        A: <input type="number" value={a} onChange={(e) => setA(e.target.value)} />
        B: <input type="number" value={b} onChange={(e) => setB(e.target.value)} />
        <button type="submit">Add</button>
      </form>

      {result && <h3>Result: {result.A} + {result.B} = {result.Sum}</h3>}
      {error && <p style={{ color: "red" }}>{error}</p>}
    </div>
  );
}

export default App;
