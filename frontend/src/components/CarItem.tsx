import { useState } from "react";
import type { User } from "../types/auth";
import type { Car } from "../types/cars";
import { useNavigate } from "react-router-dom";

export default function CarItem({ car, user, onUpdate, openModal }: { car: Car, user: User | null, onUpdate: () => void, openModal: (car: Car) => void }) {
    
    const navigate = useNavigate();
    
    const handleRentInitiation = () => {
        if (!user) {
            alert("Пожалуйста, войдите в систему, чтобы арендовать автомобиль.");
            return;
        }
        openModal(car);
    };

    return (
        <>
            <div className="border rounded-lg p-4 shadow-lg hover:shadow-xl transition-shadow">
                <img src={car.image_url} alt={car.name} className="w-full h-72 object-cover rounded-md mb-4" />
                <h3 className="text-xl font-semibold text-gray-700">{car.name}</h3>
                <p className="mb-1">{car.year}</p>
                <div className="flex justify-between">
                    <div className="flex gap-5">
                        {car.available ? (
                            <button disabled={!user} onClick={handleRentInitiation} className="bg-green-600 text-white px-6 py-2 rounded hover:bg-green-700">{user ? "В аренду" : "Войдите для аренды"}</button>
                        ) : (<button disabled={true} className="bg-orange-600 text-white px-6 py-2 rounded">В аренде</button>)}
                        <button className="bg-blue-600 text-white px-6 py-2 rounded hover:bg-blue-700" onClick={() => navigate(`/cars/${car.id}`)}>Отзывы</button>
                    </div>
                    <p className="text-base">Цена: <span className="font-semibold">{car.price_per_hour}₽/час</span></p>
                </div>
            </div>
        </>

    )

}