import React, { useState, useEffect } from 'react';
import { FiEdit, FiTrash2, FiPlus, FiSearch, FiChevronLeft, FiChevronRight } from 'react-icons/fi';
import { getOffices, createOffice, deleteOffice, updateOffice } from '../../services/offices';
import { getProvinces,getRegenciesByProvince, getDistrictsByRegency } from '../../services/locations';
import { CreateUpdateOffice, Office as OfficeType } from '../../types/office';
import { MetaData } from '../../types/pagination';
import { BadRequestError } from '../../types/error';
import DeleteModal, { DeleteModalData } from '../../components/modal/DeleteModal';
import { Province, Regency, District } from '../../types/location';
import { SelectOption } from '../../types/utils';
import MasterOfficeModal from '../../components/modal/MasterOfficeModal';

const Office: React.FC = () => {
  const moduleName = 'Kantor';
  const [offices, setOffices] = useState<OfficeType[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [searchTerm, setSearchTerm] = useState<string>('');
  const [metadata, setMetadata] = useState<MetaData>({
    current_page: 1,
    per_page: 10,
    total_items: 0
  });

  // Province state
  const [provinces, setProvinces] = useState<Province[]>([]);
  const [loadingProvinces, setLoadingProvinces] = useState<boolean>(false);
  const [selectedProvince, setSelectedProvince] = useState<Province | null>(null);
  const [provinceSearchTerm, setProvinceSearchTerm] = useState<string>('');
  const filteredProvinces = provinceSearchTerm === '' 
    ? provinces 
    : provinces.filter((province) => 
        province.name.toLowerCase().includes(provinceSearchTerm.toLowerCase())
      );

   // Regency state
   const [regencies, setRegencies] = useState<Regency[]>([]);
   const [loadingRegencies, setLoadingRegencies] = useState<boolean>(false);
   const [selectedRegency, setSelectedRegency] = useState<Regency | null>(null);
   const [regencySearchTerm, setRegencySearchTerm] = useState<string>('');
   const filteredRegencies = regencySearchTerm === '' 
     ? regencies 
     : regencies.filter((regency) => 
         regency.name.toLowerCase().includes(regencySearchTerm.toLowerCase())
       );

   // District state
   const [districts, setDistricts] = useState<District[]>([]);
   const [loadingDistricts, setLoadingDistricts] = useState<boolean>(false);
   const [selectedDistrict, setSelectedDistrict] = useState<District | null>(null);
   const [districtSearchTerm, setDistrictSearchTerm] = useState<string>('');
   const filteredDistricts = districtSearchTerm === '' 
     ? districts 
     : districts.filter((district) => 
         district.name.toLowerCase().includes(districtSearchTerm.toLowerCase())
       );

  // Modal state
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
  const [isModalAnimating, setIsModalAnimating] = useState(false);

  // Delete confirmation modal state
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState<boolean>(false);
  const [isDeleteModalAnimating, setIsDeleteModalAnimating] = useState(false);
  const [officeToDelete, setOfficeToDelete] = useState<DeleteModalData | null>(null);
  const [isDeleting, setIsDeleting] = useState<boolean>(false);

  const openModal = () => {
    setIsModalOpen(true);
    setIsModalAnimating(false); 
  };

  const closeModal = () => {
    // First start the closing animation
    setIsModalAnimating(true);
    setFormData({
      code: '',
      name: '',
      province_id: 0,
      regency_id: 0,
      district_id: 0,
      email: '',
      phone: '',
      address: '',
      logo_url: ''
    });
    
    setSelectedProvince(null);
    setSelectedRegency(null);
    setSelectedDistrict(null);

    // Then remove the modal from DOM after animation completes
    setTimeout(() => {
      setIsModalOpen(false);
      setIsModalAnimating(false);
    }, 300);
  };

  const [isEditModalOpen, setIsEditModalOpen] = useState<boolean>(false);
  const [isEditModalAnimating, setIsEditModalAnimating] = useState(false);

  const openEditModal = (office: OfficeType) => {
    setIsEditModalOpen(true);
    setIsEditModalAnimating(false); 
    setFormData({
      id: office.id,
      code: office.code,
      name: office.name,
      province_id: office.province_id,
      regency_id: office.regency_id,
      district_id: office.district_id,
      email: office.email,
      phone: office.phone,
      address: office.address,
      logo_url: office.logo_url
    });
    setSelectedProvince({
      id: office.province_id,
      name: office.province
    });
    fetchRegencies(office.province_id);
    setSelectedRegency({
      id: office.regency_id,
      name: office.regency
    });
    fetchDistricts(office.regency_id);
    setSelectedDistrict({
      id: office.district_id,
      name: office.district
    });
  }

  const closeEditModal = () => {
    // First start the closing animation
    setIsEditModalAnimating(true);
    setFormData({
      id: 0,
      code: '',
      name: '',
      province_id: 0,
      regency_id: 0,
      district_id: 0,
      email: '',
      phone: '',
      address: '',
      logo_url: ''
    });
    
    setSelectedProvince(null);
    setSelectedRegency(null);
    setSelectedDistrict(null);
    
    // Then remove the modal from DOM after animation completes
    setTimeout(() => {
      setIsEditModalOpen(false);
      setIsEditModalAnimating(false);
    }, 300);
  }

  const [formData, setFormData] = useState<CreateUpdateOffice>({
    id: 0,
    code: '',
    name: '',
    province_id: 0,
    regency_id: 0,
    district_id: 0,
    email: '',
    phone: '',
    address: '',
    logo_url: ''
  });
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const [formErrors, setFormErrors] = useState<{[key: string]: string}>({});

  // Fetch offices from API
  const fetchOffices = async (page: number = 1) => {
    setLoading(true);

    setTimeout(async() => {
      try {
        // Replace with your actual API endpoint
        const response = await getOffices(page, metadata.per_page);
        
        setOffices(response.data);
        setMetadata(response.meta_data);
      } catch (error) {
        console.error('Error fetching offices:', error);
      } finally {
        setLoading(false);
      }
    }, 500);
  };

  useEffect(() => {
    fetchOffices();
  }, []);

  // Fetch provinces from API
  const fetchProvinces = async (searchTerm: string = '') => {
    setLoadingProvinces(true);
    setTimeout(async() => {
      try {
        const response = await getProvinces(searchTerm);
        if (response.data.length > 0) {
          setProvinces(response.data);
        }
      } catch (error) {
        console.error('Error fetching provinces:', error);
      } finally {
        setLoadingProvinces(false);
      }
    }, 500);
  };

  // Fetch provinces from API
  const fetchRegencies = async (provinceId: number, searchTerm: string = '') => {
    setLoadingRegencies(true);
    setTimeout(async() => {
      try {
        const response = await getRegenciesByProvince(provinceId, searchTerm);
        if (response.data.length > 0) {
          setRegencies(response.data);
        }
      } catch (error) {
        console.error('Error fetching regencies:', error);
      } finally {
        setLoadingRegencies(false);
      }
    }, 500);
  };

  // Fetch provinces from API
  const fetchDistricts = async (regencyId: number, searchTerm: string = '') => {
    setLoadingDistricts(true);
    setTimeout(async() => {
      try {
        const response = await getDistrictsByRegency(regencyId, searchTerm);
        if (response.data.length > 0) {
          setDistricts(response.data);
        }
      } catch (error) {
        console.error('Error fetching districts:', error);
      } finally {
        setLoadingDistricts(false);
      }
    }, 500);
  };

  // Fetch provinces when modal opens
  useEffect(() => {
    if (isModalOpen) {
      fetchProvinces(provinceSearchTerm);
    }
  }, [isModalOpen]);

  // Handle province selection
  const handleProvinceChange = (province: SelectOption) => {
    setSelectedProvince(province);
    setSelectedRegency(null);
    setSelectedDistrict(null);
    fetchRegencies(province.id);
    setFormData(prev => ({
      ...prev,
      province_id: province.id,
      province: province.name
    }));
    
    // Clear error for province field
    if (formErrors.province) {
      setFormErrors(prev => {
        const newErrors = {...prev};
        delete newErrors.province;
        return newErrors;
      });
    }
  };

  const handleRegencyChange = (regency: SelectOption) => {
    setSelectedRegency(regency);
    setSelectedDistrict(null);
    fetchDistricts(regency.id);
    setFormData(prev => ({
      ...prev,
      regency_id: regency.id,
      regency: regency.name
    }));
    
    // Clear error for regency field
    if (formErrors.regency) {
      setFormErrors(prev => {
        const newErrors = {...prev};
        delete newErrors.regency;
        return newErrors;
      });
    }
  };

  const handleDistrictChange = (district: SelectOption) => {
    setSelectedDistrict(district);
    setFormData(prev => ({
      ...prev,
      district_id: district.id,
      district: district.name
    }));
    
    // Clear error for district field
    if (formErrors.district) {
      setFormErrors(prev => {
        const newErrors = {...prev};
        delete newErrors.district;
        return newErrors;
      });
    }
  };

  const handleProvinceQueryChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setProvinceSearchTerm(e.target.value);
    fetchProvinces(e.target.value);
  };

  const handleRegencyQueryChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setRegencySearchTerm(e.target.value);
    if (selectedProvince) {
      fetchRegencies(selectedProvince.id, e.target.value);
    }
  };

  const handleDistrictQueryChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setDistrictSearchTerm(e.target.value);
    if (selectedRegency) {
      fetchDistricts(selectedRegency.id, e.target.value);
    }
  };

  const locationData = {
    province: {
      data: {
        values: provinces,
        label: "Provinsi",
        isLoading: loadingProvinces,
        selectedValue: selectedProvince,
        searchTerm: provinceSearchTerm,
        filteredOptions: filteredProvinces,
        fieldError: formErrors.province_id,
        placeholder: "Pilih Provinsi"
      },
      onChange: handleProvinceChange,
      handleQueryChange: handleProvinceQueryChange,
      handleClearSearch: () => {
        setProvinceSearchTerm('');
        fetchProvinces();
      }
    },
    regency: {
      data: {
        values: regencies,
        label: "Kabupaten/Kota",
        isLoading: loadingRegencies,
        selectedValue: selectedRegency,
        searchTerm: regencySearchTerm,
        filteredOptions: filteredRegencies,
        fieldError: formErrors.regency_id,
        placeholder: "Pilih Kabupaten/Kota"
      },
      onChange: handleRegencyChange,
      handleQueryChange: handleRegencyQueryChange,
      handleClearSearch: () => {
        setRegencySearchTerm('');
        if (selectedProvince) {
          fetchRegencies(selectedProvince.id);
        }
      }
    },
    district: {
      data: {
        values: districts,
        label: "Kecamatan",
        isLoading: loadingDistricts,
        selectedValue: selectedDistrict,
        searchTerm: districtSearchTerm,
        filteredOptions: filteredDistricts,
        fieldError: formErrors.district_id,
        placeholder: "Pilih Kecamatan"
      },
      onChange: handleDistrictChange,
      handleQueryChange: handleDistrictQueryChange,
      handleClearSearch: () => {
        setDistrictSearchTerm('');
        if (selectedRegency) {
          fetchDistricts(selectedRegency.id);
        }
      }
    }
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
    
    // Clear error for this field when user starts typing
    if (formErrors[name]) {
      setFormErrors(prev => {
        const newErrors = {...prev};
        delete newErrors[name];
        return newErrors;
      });
    }
  };

  // Validate form
  const validateForm = () => {
    const errors: {[key: string]: string} = {};
    
    if (!formData.code.trim()) errors.code = "Kode kantor wajib diisi";
    if (!formData.name.trim()) errors.name = "Nama kantor wajib diisi";
    if (!formData.province_id) errors.province_id = "Provinsi wajib diisi";
    if (!formData.regency_id) errors.regency_id = "Kabupaten/Kota wajib diisi";
    if (!formData.district_id) errors.district_id = "Kecamatan wajib diisi";
    
    // Email validation
    if (formData.email && !/^\S+@\S+\.\S+$/.test(formData.email)) {
      errors.email = "Format email tidak valid";
    }
    
    setFormErrors(errors);
    return Object.keys(errors).length === 0;
  };

  // Submit new office
  const handleSubmit = async () => {
    if (!validateForm()) return;
    
    setIsSubmitting(true);
    try {
      const response = await createOffice(formData);
      
      // Add the new office to the list and close modal
      setOffices(prev => [response.data, ...prev]);
      closeModal();
      
      // Refresh the list to ensure correct pagination
      fetchOffices(metadata.current_page);
    } catch (error) {
      console.error('Error creating office:', error);
      // Handle error, maybe show an error message

      if (error instanceof BadRequestError) {
        error.fields.forEach((e) => {
          if (e === "code") {
            setFormErrors(prev => ({
              ...prev,
              code: error.message
            }));
          } else if (e === "email") {
            setFormErrors(prev => ({
              ...prev,
              email: error.message
            }));
          }
        })
      }
    } finally {
      setIsSubmitting(false);
    }
  };

  // Submit new office
  const handleEditSubmit = async () => {
    if (!validateForm()) return;
    
    setIsSubmitting(true);
    try {
      if (!formData.id) {
        setIsSubmitting(false);
        console.error('Office ID is required for update');
        return;
      }
      const response = await updateOffice(formData.id, formData);
      
      // Add the new office to the list and close modal
      setOffices(prev => [response.data, ...prev]);
      closeEditModal();
      
      // Refresh the list to ensure correct pagination
      fetchOffices(metadata.current_page);
    } catch (error) {
      console.error('Error creating office:', error);
      // Handle error, maybe show an error message

      if (error instanceof BadRequestError) {
        error.fields.forEach((e) => {
          if (e === "code") {
            setFormErrors(prev => ({
              ...prev,
              code: error.message
            }));
          } else if (e === "email") {
            setFormErrors(prev => ({
              ...prev,
              email: error.message
            }));
          }
        })
      }
    } finally {
      setIsSubmitting(false);
    }
  };

  const handlePageChange = (newPage: number) => {
    if (newPage > 0 && newPage <= Math.ceil(metadata.total_items / metadata.per_page)) {
      fetchOffices(newPage);
    }
  };

  const filteredOffices = offices.filter(office => 
    office.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    office.code.toLowerCase().includes(searchTerm.toLowerCase()) ||
    office.province.toLowerCase().includes(searchTerm.toLowerCase()) ||
    office.regency.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const totalPages = Math.ceil(metadata.total_items / metadata.per_page);

  // Open delete confirmation modal
  const openDeleteModal = (office: OfficeType) => {
    setOfficeToDelete({
      id: office.id,
      label: moduleName,
      name: office.name
    });
    setIsDeleteModalOpen(true);
    setIsDeleteModalAnimating(false);
  };

  // Close delete confirmation modal
  const closeDeleteModal = () => {
    setIsDeleteModalAnimating(true);
    
    setTimeout(() => {
      setIsDeleteModalOpen(false);
      setIsDeleteModalAnimating(false);
      setOfficeToDelete(null);
    }, 300);
  };

  // Handle delete office
  const handleDelete = async () => {
    if (!officeToDelete) return;
    
    setIsDeleting(true);
    
    try {
      await deleteOffice(Number(officeToDelete.id));
      // Refresh the offices list
      fetchOffices(metadata.current_page);
      closeDeleteModal();
    } catch (error) {
      console.error('Failed to delete office:', error);
    } finally {
      setIsDeleting(false);
    }
  };

  return (
    <div className="container mx-auto">
      <div className="flex justify-between items-center mb-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Kantor</h1>
          <p className="mt-1 text-sm text-gray-500">
          Kelola semua lokasi kantor dalam sistem
          </p>
        </div>
        <button 
          className="bg-amber-500 hover:bg-amber-600 text-white px-4 py-2 rounded-md flex items-center"
          onClick={openModal}
        >
          <FiPlus className="mr-2" />
          Tambah Kantor
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
                  #
                </th>
                <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Kode
                </th>
                <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Nama
                </th>
                <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Lokasi
                </th>
                <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Kontak
                </th>
                <th scope="col" className="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Aksi
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
                filteredOffices.map((office, index) => (
                  <tr key={office.id} className="hover:bg-gray-50">
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                      #{((metadata.current_page - 1) * metadata.per_page) + index + 1}
                    </td>
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
                        <button className="text-amber-600 hover:text-amber-900"
                          onClick={() => openEditModal(office)}>
                          <div className="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-amber-100 sm:mx-0 sm:h-8 sm:w-8">
                            <FiEdit size={18} />
                          </div>
                        </button>
                        <button className="text-red-600 hover:text-red-900"
                          onClick={() => openDeleteModal(office)}>
                          <div className="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-red-100 sm:mx-0 sm:h-8 sm:w-8">
                            <FiTrash2 size={18} />
                          </div>
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
              onClick={() => handlePageChange(metadata.current_page - 1)}
              disabled={metadata.current_page === 1}
              className={`relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md ${
                metadata.current_page === 1 
                  ? 'bg-gray-100 text-gray-400 cursor-not-allowed' 
                  : 'bg-white text-gray-700 hover:bg-gray-50'
              }`}
            >
              Previous
            </button>
            <button
              onClick={() => handlePageChange(metadata.current_page + 1)}
              disabled={metadata.current_page === totalPages}
              className={`relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md ${
                metadata.current_page === totalPages 
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
                Showing <span className="font-medium">{(metadata.current_page - 1) * metadata.per_page + 1}</span> to{' '}
                <span className="font-medium">
                  {Math.min(metadata.current_page * metadata.per_page, metadata.total_items)}
                </span>{' '}
                of <span className="font-medium">{metadata.total_items}</span> results
              </p>
            </div>
            <div>
              <nav className="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
                <button
                  onClick={() => handlePageChange(metadata.current_page - 1)}
                  disabled={metadata.current_page === 1}
                  className={`relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium ${
                    metadata.current_page === 1 
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
                    Math.abs(page - metadata.current_page) < 2
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
                          metadata.current_page === page
                            ? 'z-10 bg-amber-50 border-amber-500 text-amber-600'
                            : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50'
                        }`}
                      >
                        {page}
                      </button>
                    );
                  })}
                
                <button
                  onClick={() => handlePageChange(metadata.current_page + 1)}
                  disabled={metadata.current_page === totalPages}
                  className={`relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium ${
                    metadata.current_page === totalPages 
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

      {/* Modal for adding new office */}
      {isModalOpen && (
        <MasterOfficeModal
          formTitle="Tambah Kantor Baru"
          isSubmitting={isSubmitting}
          isModalAnimating={isModalAnimating}
          formData={formData}
          formErrors={formErrors}
          locationData={locationData}
          closeModal={closeModal}
          onAnimationEnd={() => setIsModalAnimating(false)}
          handleInputChange={handleInputChange}
          handleSubmit={handleSubmit}
        />
      )}

      {isEditModalOpen && (
        <MasterOfficeModal
          formTitle="Ubah Kantor"
          isSubmitting={isSubmitting}
          isModalAnimating={isEditModalAnimating}
          formData={formData}
          formErrors={formErrors}
          locationData={locationData}
          closeModal={closeEditModal}
          onAnimationEnd={() => setIsEditModalAnimating(false)}
          handleInputChange={handleInputChange}
          handleSubmit={handleEditSubmit}
        />
      )}
      
      {/* Delete Confirmation Modal */}
      {isDeleteModalOpen && (
        <DeleteModal
          isDeleteLoading={isDeleting}
          isDeleteModalAnimating={isDeleteModalAnimating}
          deleteModalData={officeToDelete}
          handleDelete={handleDelete}
          closeDeleteModal={closeDeleteModal}
          onAnimationEnd={() => setIsDeleteModalAnimating(false)}
        />
      )}
    </div>
  );
};

export default Office;