import { useEffect, useState } from "react";
import type { Car } from "../types/cars";
import CarItem from "../components/CarItem";
import { carService } from "../services/carService";

export default function Home() {
    const [cars, setCars] = useState<Car[]>([]);

    const fetchCars = async () => {
        try {
            const cars = await carService.getCars();
            setCars(cars);
        } catch (error) {
            console.log(error)
        }
    }

    useEffect(() => {
        fetchCars();
    }, [])

    return (
        <section>
            <h2 className="text-2xl font-semibold mb-4 text-gray-800">Автомобили</h2>
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {cars.map((car) => (
                        <CarItem key={car.id} car={car}></CarItem>
                    ))}
                </div>
        </section>
    );
};