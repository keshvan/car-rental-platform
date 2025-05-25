import { Link, Outlet } from "react-router-dom";
import { useAuth } from "./AuthContext";
import { useState } from "react";
import { carService } from "../services/carService";

export default function Layout() {
    const { user, logout, setUser } = useAuth();
    const [openModal, setOpenModal] = useState(false);
    const [amount, setAmount] = useState(0);

    const handleDeposit = async () => {
        if (amount > 0 && user) {
            try {
                const newBalance = await carService.updateBalance(user.id, amount);
                setUser({ ...user, balance: user.balance + newBalance.amount });
            } catch (error) {
                console.error(error);
            }
        }
    }

    return (
        <div className="flex flex-col h-screen text-sm">
            {openModal && (
                <div className="fixed inset-0 bg-black/50 flex justify-center items-center">
                    <div className="bg-white p-4 rounded-lg">
                        <h2 className="text-lg font-bold">Пополнить баланс</h2>
                        <div className="flex gap-2">
                            <input onChange={(e) => setAmount(Number(e.target.value))} type="number" className="border rounded px-2 py-1" />
                            <button className="bg-blue-500 text-white px-4 py-2 rounded" onClick={() => { handleDeposit(); setOpenModal(false) }}>Пополнить</button>
                        </div>
                    </div>
                </div>
            )}
            <header className="bg-gray-100 border-b border-gray-300 px-4 py-2 flex justify-between">
                <Link to={"/"}>
                    <h1 className="text-lg font-bold text-gray-800">Аренда автомобилей</h1>
                </Link>
                <div className="flex gap-5 text-base items-center">
                    {user ? (
                        <>
                            <button className="bg-blue-500 text-white px-4 py-2 rounded" onClick={() => setOpenModal(true)}>Пополнить</button>
                            <div className="flex gap-5 items-center">
                                <p>{user.balance} ₽</p>
                                <Link to={"/profile"}>{user.email}</Link>
                            </div>
                            
                            <button
                                onClick={logout}
                                className="py-1 text-red-600 hover:text-red-500 hover:underline"
                            >
                                Выйти
                            </button>
                        </>
                    ) : (
                        <>
                            <Link to="/login" className="py-1 text-blue-600 hover:text-blue-500 hover:underline">Войти</Link>
                            <Link to="/register" className="py-1 text-blue-600 hover:text-blue-500 hover:underline">Регистрация</Link>
                        </>
                    )}
                </div>
            </header>
            <div className="flex-1 flex h-0 min-h-0 overflow-hidden">
                <div className="flex-1 overflow-y-auto">
                    <Outlet />
                </div>
            </div>
        </div>
    );
}