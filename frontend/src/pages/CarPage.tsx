import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { carService } from '../services/carService';
import type { Car, Review } from '../types/cars';

export default function CarPage() {
    const { id } = useParams<{ id: string }>();
    const [car, setCar] = useState<Car | null>(null);
    const [reviews, setReviews] = useState<Review[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const res = await carService.getCarWithReviews(id);
                setCar(res.car);
                setReviews(res.reviews);
            } catch (err) {
                setError(err as string);
            } finally {
                setLoading(false);
            }
        }
        fetchData();
    }, [id]);

    if (loading) {
        return <div className="container mx-auto p-4">Загрузка...</div>;
    }

    if (error) {
        return <div className="container mx-auto p-4 text-red-500">{error}</div>;
    }

    if (!car) {
        return <div className="container mx-auto p-4">Автомобиль не найден.</div>;
    }

    return (
        <div className="container mx-auto p-4">
            <div className="bg-white shadow-lg rounded-lg overflow-hidden">
                <img src={car.image_url || 'https://via.placeholder.com/600x400'} alt={car.name} className="w-full h-96 object-cover" />
                <div className="p-6">
                    <h1 className="text-3xl font-bold mb-2 text-gray-800">{car.name}</h1>
                    <p className="text-xl text-gray-600 mb-4">Год выпуска: {car.year}</p>

                    <div className="mt-6">
                        <h2 className="text-2xl font-semibold mb-3 text-gray-700">Отзывы</h2>
                        {(reviews && reviews.length >  0) ? (
                            <ul className="space-y-4">
                                {reviews.map((review) => (
                                    <li key={review.id} className="p-4 border rounded-md bg-gray-50">
                                        <div className="flex justify-between items-center mb-1">
                                            <span className="font-semibold text-gray-700">{review.email}</span>
                                            <span className="text-sm text-gray-500">Оценка: {"⭐".repeat(review.rating)}</span>
                                        </div>
                                        <p className="text-gray-600">{review.comment}</p>
                                        <p className="text-xs text-gray-400 mt-1">{new Date(review.created_at).toLocaleString()}</p>
                                    </li>
                                ))}
                            </ul>
                        ) : (
                            <p className="text-gray-600">Отзывов пока нет.</p>
                        )}
                    </div>
                </div>
            </div>
        </div>
    );
}