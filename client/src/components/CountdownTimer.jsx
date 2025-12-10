import { useState, useEffect } from 'react';

function CountdownTimer({ createdAt, status, onExpired }) {
  const [timeLeft, setTimeLeft] = useState(null);
  const [isExpired, setIsExpired] = useState(false);

  useEffect(() => {
    if (status !== 'pending') {
      setTimeLeft(null);
      return;
    }

    const calculateTimeLeft = () => {
      const created = new Date(createdAt);
      const expiryTime = new Date(created.getTime() + 5 * 60 * 1000);
      const now = new Date();
      const difference = expiryTime - now;

      if (difference <= 0) {
        setIsExpired(true);
        setTimeLeft({ minutes: 0, seconds: 0 });
        if (onExpired && !isExpired) {
          onExpired();
        }
        return null;
      }

      const minutes = Math.floor((difference / 1000 / 60) % 60);
      const seconds = Math.floor((difference / 1000) % 60);

      return { minutes, seconds };
    };

    const initial = calculateTimeLeft();
    setTimeLeft(initial);

    const timer = setInterval(() => {
      const newTime = calculateTimeLeft();
      setTimeLeft(newTime);
    }, 1000);

    return () => clearInterval(timer);
  }, [createdAt, status, onExpired, isExpired]);

  if (status !== 'pending' || !timeLeft) {
    return <span className="text-gray-400 text-sm">-</span>;
  }

  const isUrgent = timeLeft.minutes < 2;

  return (
    <div className={`text-sm font-medium ${isUrgent ? 'text-red-600' : 'text-orange-600'}`}>
      {isExpired ? (
        <span className="text-red-600 font-bold">EXPIRED</span>
      ) : (
        <>
          {String(timeLeft.minutes).padStart(2, '0')}:{String(timeLeft.seconds).padStart(2, '0')}
          {isUrgent && <span className="ml-1 animate-pulse">⚠️</span>}
        </>
      )}
    </div>
  );
}

export default CountdownTimer;
