import type { Car } from "../types/cars";

export default function CarItem({ car }: { car: Car }) {
    return (
        <div className="border rounded-lg p-4 shadow-lg hover:shadow-xl transition-shadow">
            <img src={car.image_url} alt={car.name} className="w-full h-48 object-cover rounded-md mb-4" />
            <h3 className="text-xl font-semibold text-gray-700">{car.name}</h3>
            <p>{car.year}</p>
            <p className="text-gray-600">Цена: {car.price_per_hour}₽/час</p>
            {car.available && (
                <button className="bg-green-600 text-white px-6 py-2 rounded hover:bg-green-700">В аренду</button>
            )}
        </div>
    )

}