import { useEffect, useState, type ChangeEvent } from "react";
import type { Brand, Car } from "../types/cars";
import CarItem from "../components/CarItem";
import { carService } from "../services/carService";
import { useAuth } from "../components/AuthContext";

export default function Home() {
    const [cars, setCars] = useState<Car[]>([]);
    const [brands, setBrands] = useState<Brand[]>([]);

    const [available, setAvailable] = useState(true);
    const [brand, setBrand] = useState<number | "all">("all");
    const [search, setSearch] = useState("");
    const { user } = useAuth();
    const [openModal, setOpenModal] = useState(false);
    const [carToRent, setCarToRent] = useState<Car | null>(null);

    const fetchCars = async () => {
        try {
            const carsData = await carService.getCars();
            setCars(carsData);
        } catch (error) {
            console.error(error);
        }
    }

    const fetchBrands = async () => {
        try {
            const brandsData = await carService.getBrands();
            setBrands(brandsData);
        } catch (error) {
            console.error(error);
        }
    }

    const handleOpenModal = (car: Car) => {
        setCarToRent(car);
        setOpenModal(true);
    }

    const handleRent = async (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
        if (carToRent && user) {
            try {
                await carService.newRent(carToRent.id, new Date().toISOString());
            } catch (error) {
                console.error(error);
            }
        }
        setOpenModal(false);
        setCarToRent(null);
    }

    useEffect(() => {
        fetchCars();
        fetchBrands();
    }, []);

    const filteredCars = cars.filter(car => {
        if (available && !car.available) return false;
        if (brand !== "all" && car.brand_id !== brand) return false;
        if (search.trim() !== "" && !car.name.toLowerCase().includes(search.toLowerCase())) return false;
        return true;
    });

    const onBrandChange = (e: ChangeEvent<HTMLSelectElement>) => {
        const val = e.target.value;
        setBrand(val === "all" ? "all" : Number(val));
    };

    return (
        <section className="py-5 px-5">
            {openModal && carToRent && (
                <div className="fixed inset-0 bg-black/50 flex justify-center items-center">
                    <div className="bg-white p-4 rounded-lg">
                        <h2 className="text-lg font-bold">Вы точно хотите арендовать автомобиль: {carToRent.name}?</h2>
                        <div className="flex gap-2">
                            <button className="bg-red-500 text-white px-4 py-2 rounded" onClick={() => { setOpenModal(false); setCarToRent(null); }}>Отменить</button>
                            <button className="bg-blue-500 text-white px-4 py-2 rounded" onClick={(e) => {handleRent(e)}}>Подтвердить</button>
                        </div>
                    </div>
                </div>
            )}
            <h2 className="text-2xl font-semibold mb-4 text-gray-800">Автомобили</h2>

            <div className="flex items-center mb-6 space-x-6">
                <div>
                    <input
                        type="text"
                        placeholder="Поиск по названию..."
                        value={search}
                        onChange={(e) => setSearch(e.target.value)}
                        className="w-full border rounded px-3 py-2"
                    />
                </div>
                <label className="flex items-center space-x-2">
                    <input
                        type="checkbox"
                        checked={available}
                        onChange={(e) => setAvailable(e.target.checked)}
                        className="form-checkbox"
                    />
                    <span>Только доступные</span>
                </label>

                <div>
                    <label htmlFor="brandSelect" className="mr-2">
                        Бренд:
                    </label>
                    <select
                        id="brandSelect"
                        value={brand}
                        onChange={onBrandChange}
                        className="border rounded px-2 py-1"
                    >
                        <option value="all">Все бренды</option>
                        {brands.map((b) => (
                            <option key={b.id} value={b.id}>
                                {b.name}
                            </option>
                        ))}
                    </select>
                </div>
            </div>
            <div className="grid grid-cols-3 gap-6">
                {filteredCars.map((car) => (
                    <CarItem key={car.id} car={car} user={user} onUpdate={fetchCars} openModal={handleOpenModal}></CarItem>
                ))}
            </div>
        </section>
    );
};