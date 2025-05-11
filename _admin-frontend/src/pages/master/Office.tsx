import React, { useState, useEffect } from 'react';
import { FiEdit, FiTrash2, FiPlus, FiSearch, FiChevronLeft, FiChevronRight } from 'react-icons/fi';

interface Office {
  id: string;
  index: string;
  code: string;
  name: string;
  province: string;
  regency: string;
  district: string;
  email: string;
  phone: string;
  address: string;
  logo_url: string;
  created_by: string;
}

interface MetaData {
  currentPage: number;
  perPage: number;
  totalItems: number;
}

const Office: React.FC = () => {
  const [offices, setOffices] = useState<Office[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [searchTerm, setSearchTerm] = useState<string>('');
  const [metadata, setMetadata] = useState<MetaData>({
    currentPage: 1,
    perPage: 10,
    totalItems: 0
  });

  // Fetch offices from API
  const fetchOffices = async (page: number = 1) => {
    setLoading(true);
    try {
      // Replace with your actual API endpoint
      const response = await fetch(`/api/offices?page=${page}&limit=${metadata.perPage}`);
      const data = await response.json();
      
      if (data.data) {
        setOffices(data.data);
        setMetadata(data.metaData);
      }
    } catch (error) {
      console.error('Error fetching offices:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchOffices();
  }, []);

  // For demo purposes only - remove this in production
  useEffect(() => {
    // Mock data for development
    const mockOffices: Office[] = Array.from({ length: 10 }, (_, i) => ({
      id: `${i + 1}`,
      index: `#${i + 1}`,
      code: `OFC${i + 100}`,
      name: `Office Location ${i + 1}`,
      province: `Province ${i % 3 + 1}`,
      regency: `Regency ${i % 5 + 1}`,
      district: `District ${i % 7 + 1}`,
      email: `office${i + 1}@example.com`,
      phone: `+1234567890${i}`,
      address: `123 Main St, Building ${i + 1}`,
      logo_url: '',
      created_by: '1'
    }));

    setOffices(mockOffices);
    setMetadata({
      currentPage: 1,
      perPage: 10,
      totalItems: 50
    });
    setLoading(false);
  }, []);

  const handlePageChange = (newPage: number) => {
    if (newPage > 0 && newPage <= Math.ceil(metadata.totalItems / metadata.perPage)) {
      fetchOffices(newPage);
    }
  };

  const filteredOffices = offices.filter(office => 
    office.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    office.code.toLowerCase().includes(searchTerm.toLowerCase()) ||
    office.province.toLowerCase().includes(searchTerm.toLowerCase()) ||
    office.regency.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const totalPages = Math.ceil(metadata.totalItems / metadata.perPage);

  return (
    <div className="container mx-auto">
      <div className="flex justify-between items-center mb-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Offices</h1>
          <p className="mt-1 text-sm text-gray-500">
            Manage all office locations in the system
          </p>
        </div>
        <button 
          className="bg-amber-500 hover:bg-amber-600 text-white px-4 py-2 rounded-md flex items-center"
          onClick={() => {/* Handle new office creation */}}
        >
          <FiPlus className="mr-2" />
          Add Office
        </button>
      </div>
      
      <div className="bg-white shadow-md rounded-lg overflow-hidden">
        {/* Search and filters */}
        <div className="p-4 border-b border-gray-200">
          <div className="relative">
            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <FiSearch className="text-gray-400" />
            </div>
            <input
              type="text"
              className="pl-10 pr-4 py-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring-2 focus:ring-amber-500 focus:border-transparent"
              placeholder="Search offices..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
          </div>
        </div>
        
        {/* Table */}
        <div className="overflow-x-auto">
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Code
                </th>
                <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Name
                </th>
                <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Location
                </th>
                <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Contact
                </th>
                <th scope="col" className="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Actions
                </th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {loading ? (
                <tr>
                  <td colSpan={5} className="px-6 py-4 text-center">
                    <div className="flex justify-center">
                      <div className="animate-spin rounded-full h-6 w-6 border-b-2 border-amber-500"></div>
                    </div>
                  </td>
                </tr>
              ) : filteredOffices.length === 0 ? (
                <tr>
                  <td colSpan={5} className="px-6 py-4 text-center text-sm text-gray-500">
                    No offices found
                  </td>
                </tr>
              ) : (
                filteredOffices.map((office) => (
                  <tr key={office.id} className="hover:bg-gray-50">
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                      {office.code}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      {office.name}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      <div className="flex flex-col">
                        <span>{office.province}</span>
                        <span className="text-xs text-gray-400">{office.regency}, {office.district}</span>
                      </div>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      <div className="flex flex-col">
                        <span>{office.email}</span>
                        {office.phone && <span className="text-xs text-gray-400">{office.phone}</span>}
                      </div>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                      <div className="flex justify-end space-x-2">
                        <button 
                          className="text-amber-600 hover:text-amber-900"
                          onClick={() => {/* Handle edit */}}
                        >
                          <FiEdit size={18} />
                        </button>
                        <button 
                          className="text-red-600 hover:text-red-900"
                          onClick={() => {/* Handle delete */}}
                        >
                          <FiTrash2 size={18} />
                        </button>
                      </div>
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>
        
        {/* Pagination */}
        <div className="px-6 py-3 flex items-center justify-between border-t border-gray-200">
          <div className="flex-1 flex justify-between sm:hidden">
            <button
              onClick={() => handlePageChange(metadata.currentPage - 1)}
              disabled={metadata.currentPage === 1}
              className={`relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md ${
                metadata.currentPage === 1 
                  ? 'bg-gray-100 text-gray-400 cursor-not-allowed' 
                  : 'bg-white text-gray-700 hover:bg-gray-50'
              }`}
            >
              Previous
            </button>
            <button
              onClick={() => handlePageChange(metadata.currentPage + 1)}
              disabled={metadata.currentPage === totalPages}
              className={`relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md ${
                metadata.currentPage === totalPages 
                  ? 'bg-gray-100 text-gray-400 cursor-not-allowed' 
                  : 'bg-white text-gray-700 hover:bg-gray-50'
              }`}
            >
              Next
            </button>
          </div>
          <div className="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
            <div>
              <p className="text-sm text-gray-700">
                Showing <span className="font-medium">{(metadata.currentPage - 1) * metadata.perPage + 1}</span> to{' '}
                <span className="font-medium">
                  {Math.min(metadata.currentPage * metadata.perPage, metadata.totalItems)}
                </span>{' '}
                of <span className="font-medium">{metadata.totalItems}</span> results
              </p>
            </div>
            <div>
              <nav className="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
                <button
                  onClick={() => handlePageChange(metadata.currentPage - 1)}
                  disabled={metadata.currentPage === 1}
                  className={`relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium ${
                    metadata.currentPage === 1 
                      ? 'text-gray-300 cursor-not-allowed' 
                      : 'text-gray-500 hover:bg-gray-50'
                  }`}
                >
                  <span className="sr-only">Previous</span>
                  <FiChevronLeft className="h-5 w-5" />
                </button>
                
                {/* Page numbers */}
                {Array.from({ length: totalPages }, (_, i) => i + 1)
                  .filter(page => 
                    page === 1 || 
                    page === totalPages || 
                    Math.abs(page - metadata.currentPage) < 2
                  )
                  .map((page, i, array) => {
                    // Add ellipsis
                    if (i > 0 && array[i - 1] !== page - 1) {
                      return (
                        <span
                          key={`ellipsis-${page}`}
                          className="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700"
                        >
                          ...
                        </span>
                      );
                    }
                    
                    return (
                      <button
                        key={page}
                        onClick={() => handlePageChange(page)}
                        className={`relative inline-flex items-center px-4 py-2 border text-sm font-medium ${
                          metadata.currentPage === page
                            ? 'z-10 bg-amber-50 border-amber-500 text-amber-600'
                            : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50'
                        }`}
                      >
                        {page}
                      </button>
                    );
                  })}
                
                <button
                  onClick={() => handlePageChange(metadata.currentPage + 1)}
                  disabled={metadata.currentPage === totalPages}
                  className={`relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium ${
                    metadata.currentPage === totalPages 
                      ? 'text-gray-300 cursor-not-allowed' 
                      : 'text-gray-500 hover:bg-gray-50'
                  }`}
                >
                  <span className="sr-only">Next</span>
                  <FiChevronRight className="h-5 w-5" />
                </button>
              </nav>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Office;