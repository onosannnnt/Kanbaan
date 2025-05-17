import { useState, useEffect } from "react";

const Loading = () => {
  const [dots, setDots] = useState(1);

  useEffect(() => {
    const interval = setInterval(() => {
      setDots((prevDots) => (prevDots >= 3 ? 1 : prevDots + 1));
    }, 500);

    return () => clearInterval(interval);
  }, []);

  const renderDots = () => {
    return ". ".repeat(dots);
  };

  return (
    <div className="flex flex-col items-center justify-center w-full h-screen">
      <svg
        className="w-16 h-16 animate-spin text-gray-200"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      >
        <circle cx="12" cy="12" r="10" />
        <line x1="12" y1="2" x2="12" y2="6" />
        <line x1="12" y1="18" x2="12" y2="22" />
        <line x1="4.22" y1="4.22" x2="6.34" y2="6.34" />
        <line x1="17.66" y1="17.66" x2="19.78" y2="19.78" />
        <line x1="2" y1="12" x2="6" y2="12" />
        <line x1="18" y1="12" x2="22" y2="12" />
        <line x1="4.22" y1="19.78" x2="6.34" y2="17.66" />
        <line x1="17.66" y1="6.34" x2="19.78" y2="4.22" />
      </svg>

      <div className="mt-4 text-xl font-medium text-gray-500">
        Loading {renderDots()}
      </div>
    </div>
  );
};

export default Loading;
