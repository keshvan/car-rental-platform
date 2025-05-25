import React, { useEffect, useState } from 'react';
import { useAuth } from '../components/AuthContext';
import type { Rent } from '../types/cars';
import { carService } from '../services/carService';

export default function ProfilePage() {
  const { user } = useAuth();
  const [rents, setRents] = useState<Rent[]>([]);
  const [showModal, setShowModal] = useState(false);
  const [currentRent, setCurrentRent] = useState<Rent | null>(null);
  const [rating, setRating] = useState<number>(0);
  const [reviewText, setReviewText] = useState<string>('');

  const fetchRents = async () => {
    try {
      const res = await carService.getMyRents()
      setRents(res)
    } catch (error) {
      console.error(error)
    }
  }

  useEffect(() => {
    fetchRents()
  }, [])

  const handleOpenModal = (rent: Rent) => {
    setCurrentRent(rent);
    setShowModal(true);
    setRating(0);
    setReviewText('');
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setCurrentRent(null);
  };

  const handleSubmit = async () => {
    if (!currentRent || rating === 0) {
      alert('Пожалуйста, поставьте оценку.');
      return;
    }
    try {
      await carService.completeRent(currentRent.id, rating, reviewText);
      await fetchRents();
      handleCloseModal();
    } catch (error) {
      alert('Не удалось завершить аренду. Попробуйте снова.');
    }
  };

  if (!user) {
    return <div>Пожалуйста, войдите в систему</div>;
  }

  return (
    <div className="max-w-2xl mx-auto mt-10 p-6">
      <h2 className="text-3xl font-bold text-gray-800 mb-6">Личный кабинет</h2>

      <section className="mb-8 p-4 border rounded-lg shadow">
        <h3 className="text-xl font-semibold text-gray-700 mb-3">Информация о пользователе</h3>
        <p><strong>Email:</strong> {user.email}</p>
      </section>

      <section className="p-4 border rounded-lg shadow">
        <h3 className="text-xl font-semibold text-gray-700 mb-3">Аренды</h3>
        {rents.length > 0 ? (
          <ul className="space-y-3">
            {rents.map((rent) => (
              <li key={rent.id} className="p-3 border rounded-md bg-gray-50">
                <p><strong>Автомобиль:</strong> {rent.car_name}</p>
                <p><strong>Начало аренды:</strong> {new Date(rent.start_date).toLocaleString()}</p>
                <p><strong>Конец аренды:</strong> {rent.status === 'active' ? '–' : new Date(rent.end_date).toLocaleString()}</p>
                {rent.status === 'active' && (
                  <div>
                    <p><strong>Цена:</strong> {Math.round((new Date().getTime() - new Date(rent.start_date).getTime()) / 3600000 * rent.price_per_hour)} руб.</p>
                    <button className="bg-blue-500 text-white px-4 py-2 rounded-md" onClick={() => handleOpenModal(rent)}>
                      Завершить аренду
                    </button>
                  </div>
                )}
                {rent.status === 'completed' && (
                  <p><strong>Итого:</strong> {rent.total_price} руб.</p>
                )}
              </li>
            ))}
          </ul>
        ) : (
          <p>У вас пока нет бронирований.</p>
        )}
      </section>
      {showModal && currentRent && (
        <div className="fixed inset-0 bg-black/50 flex items-center justify-center p-4">
          <div className="bg-white p-6 rounded-lg shadow-xl w-full max-w-md">
            <h4 className="text-xl font-semibold mb-4">Завершить аренду и оставить отзыв</h4>
            <p className="mb-2">Автомобиль: <strong>{currentRent.car_name}</strong></p>

            <div className="mb-4">
              <label className="block text-sm font-medium text-gray-700 mb-1">Оценка (обязательно):</label>
              <div className="flex space-x-1">
                {[1, 2, 3, 4, 5].map((star) => (
                  <button
                    key={star}
                    onClick={() => setRating(star)}
                    className={`p-2 rounded ${rating >= star ? 'bg-yellow-400' : 'bg-gray-200'}`}
                  >
                    ⭐
                  </button>
                ))}
              </div>
            </div>

            <div className="mb-6">
              <label htmlFor="reviewText" className="block text-sm font-medium text-gray-700 mb-1">
                Отзыв (опционально):
              </label>
              <textarea
                id="reviewText"
                value={reviewText}
                onChange={(e) => setReviewText(e.target.value)}
                rows={4}
                className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                placeholder="Расскажите о вашем опыте..."
              />
            </div>

            <div className="flex justify-end space-x-3">
              <button
                onClick={handleCloseModal}
                className="px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50"
              >
                Отмена
              </button>
              <button
                onClick={handleSubmit}
                disabled={rating === 0}
                className="px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 disabled:bg-blue-300 disabled:cursor-not-allowed"
              >
                Отправить отзыв и завершить
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}