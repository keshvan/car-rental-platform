import React from 'react';

// TODO: Определить интерфейс для автомобиля (Car)
interface Car {
  id: number;
  name: string;
  imageUrl: string;
  pricePerDay: number;
}

// TODO: Заменить моковые данные на данные из API
const mockCars: Car[] = [
  {
    id: 1,
    name: 'Toyota Camry',
    imageUrl: 'https://via.placeholder.com/300x200.png?text=Toyota+Camry',
    pricePerDay: 50,
  },
  {
    id: 2,
    name: 'BMW X5',
    imageUrl: 'https://via.placeholder.com/300x200.png?text=BMW+X5',
    pricePerDay: 100,
  },
  {
    id: 3,
    name: 'Audi A6',
    imageUrl: 'https://via.placeholder.com/300x200.png?text=Audi+A6',
    pricePerDay: 90,
  },
];

interface CarListProps {
  searchTerm: string;
}

export default function CarList({ searchTerm }: CarListProps) {
  const filteredCars = mockCars.filter(car =>
    car.name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <section>
      <h2 className="text-2xl font-semibold mb-4 text-gray-800">Доступные автомобили</h2>
      {filteredCars.length > 0 ? (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {filteredCars.map((car) => (
            <div key={car.id} className="border rounded-lg p-4 shadow-lg hover:shadow-xl transition-shadow">
              <img src={car.imageUrl} alt={car.name} className="w-full h-48 object-cover rounded-md mb-4" />
              <h3 className="text-xl font-semibold text-gray-700">{car.name}</h3>
              <p className="text-gray-600">Цена: {car.pricePerDay}$/день</p>
              {/* TODO: Добавить кнопку "Подробнее" или "Забронировать" */}
            </div>
          ))}
        </div>
      ) : (
        <p className="text-lg text-gray-700 text-center">По вашему запросу автомобили не найдены.</p>
      )}
    </section>
  );
};