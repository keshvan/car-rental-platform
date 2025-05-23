import React from 'react';

const mockUser = {
  name: 'Иван Иванов',
  email: 'ivan@example.com',
  bookingHistory: [
    { id: 1, carName: 'Toyota Camry', startDate: '2024-05-01', endDate: '2024-05-05' },
    { id: 2, carName: 'BMW X5', startDate: '2024-06-10', endDate: '2024-06-15' },
  ],
};

export default function ProfilePage() {

  return (
    <div className="max-w-2xl mx-auto mt-10 p-6">
      <h2 className="text-3xl font-bold text-gray-800 mb-6">Личный кабинет</h2>

      <section className="mb-8 p-4 border rounded-lg shadow">
        <h3 className="text-xl font-semibold text-gray-700 mb-3">Информация о пользователе</h3>
        <p><strong>Имя:</strong> {mockUser.name}</p>
        <p><strong>Email:</strong> {mockUser.email}</p>
      </section>

      <section className="p-4 border rounded-lg shadow">
        <h3 className="text-xl font-semibold text-gray-700 mb-3">История бронирований</h3>
        {mockUser.bookingHistory.length > 0 ? (
          <ul className="space-y-3">
            {mockUser.bookingHistory.map((booking) => (
              <li key={booking.id} className="p-3 border rounded-md bg-gray-50">
                <p><strong>Автомобиль:</strong> {booking.carName}</p>
                <p><strong>Начало аренды:</strong> {booking.startDate}</p>
                <p><strong>Конец аренды:</strong> {booking.endDate}</p>
              </li>
            ))}
          </ul>
        ) : (
          <p>У вас пока нет бронирований.</p>
        )}
      </section>
    </div>
  );
};