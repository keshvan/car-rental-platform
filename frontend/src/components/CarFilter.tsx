import React from 'react';

interface CarFilterProps {
  searchTerm: string;
  onSearchTermChange: (term: string) => void;
}

export default function CarFilter({ searchTerm, onSearchTermChange }: CarFilterProps) {
  return (
    <section className="mb-8 p-4 border rounded-lg shadow">
      <h2 className="text-xl font-semibold mb-4 text-gray-800">Поиск и Фильтры</h2>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label htmlFor="search" className="block text-sm font-medium text-gray-700 mb-1">
            Поиск по названию
          </label>
          <input
            type="text"
            id="search"
            name="search"
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
            placeholder="Например, Toyota Camry"
            value={searchTerm}
            onChange={(e) => onSearchTermChange(e.target.value)}
          />
        </div>
      </div>
    </section>
  );
};