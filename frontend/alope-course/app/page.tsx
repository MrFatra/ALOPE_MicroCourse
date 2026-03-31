"use client";

import { useEffect, useState } from "react";
import { getCourses } from "./services/courseService";

export default function HomePage() {
  const [data, setData] = useState<any>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    getCourses()
      .then((res) => setData(res))
      .catch((err) => console.error(err))
      .finally(() => setLoading(false));
  }, []);

  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24 bg-gray-50">
      <div className="z-10 w-full max-w-5xl items-center justify-between font-mono text-sm lg:flex">
        <h1 className="text-4xl font-bold text-blue-600">Alope Course</h1>
      </div>

      <div className="mt-8 p-6 bg-white rounded-xl shadow-md">
        <h2 className="text-xl mb-4">Status API Golang:</h2>
        {loading ? (
          <p>Sedang menghubungkan ke server...</p>
        ) : (
          <pre className="bg-gray-100 p-4 rounded border">
            {JSON.stringify(data, null, 2)}
          </pre>
        )}
      </div>
    </main>
  );
}
