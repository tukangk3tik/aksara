import React, { useState, useEffect } from 'react';
import { FiEdit, FiTrash2, FiPlus, FiSearch, FiChevronLeft, FiChevronRight } from 'react-icons/fi';
import { getProvinces,getRegenciesByProvince, getDistrictsByRegency } from '../../services/locations';
import { CreateUpdateSchool, School as SchoolType } from '../../types/school';
import { MetaData } from '../../types/pagination';
import { BadRequestError } from '../../types/error';
import DeleteModal, { DeleteModalData } from '../../components/modal/DeleteModal';
import { Province, Regency, District } from '../../types/location';
import { SelectOption } from '../../types/utils';
import { createSchool, deleteSchool, getSchools, updateSchool } from '../../services/schools';
import MasterSchoolModal from '../../components/modal/MasterSchoolModal';
import { fetchOfficesSelectOption } from '../../services/offices';

const School: React.FC = () => {
  const moduleName = 'Sekolah';
  const [schools, setSchools] = useState<SchoolType[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [searchTerm, setSearchTerm] = useState<string>('');
  const [metadata, setMetadata] = useState<MetaData>({
    current_page: 1,
    per_page: 10,
    total_items: 0
  });

  // Modal state
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
  const [isModalAnimating, setIsModalAnimating] = useState(false);

  // Delete confirmation modal state
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState<boolean>(false);
  const [isDeleteModalAnimating, setIsDeleteModalAnimating] = useState(false);
  const [schoolToDelete, setSchoolToDelete] = useState<DeleteModalData | null>(null);
  const [isDeleting, setIsDeleting] = useState<boolean>(false);

  const openModal = () => {
    setIsModalOpen(true);
    setIsModalAnimating(false); 
  };

  const closeModal = () => {
    setIsModalAnimating(true);
    setFormData({
      code: '',
      name: '',
      office_id: 0,
      is_public_school: false,
      province_id: 0,
      regency_id: 0,
      district_id: 0,
      email: '',
      phone: '',
      address: '',
      logo_url: ''
    });
    
    setSelectedOffice(null);
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

  const openEditModal = (school: SchoolType) => {
    setIsEditModalOpen(true);
    setIsEditModalAnimating(false); 
    setFormData({
      id: school.id,
      code: school.code,
      name: school.name,
      office_id: school.office_id,
      is_public_school: school.is_public_school,
      province_id: school.province_id,
      regency_id: school.regency_id,
      district_id: school.district_id,
      email: school.email,
      phone: school.phone,
      address: school.address,
      logo_url: school.logo_url
    });
    setSelectedOffice({
      id: school.office_id,
      name: school.office,
    });
    setSelectedProvince({
      id: school.province_id,
      name: school.province
    });
    fetchRegencies(school.province_id);
    setSelectedRegency({
      id: school.regency_id,
      name: school.regency
    });
    fetchDistricts(school.regency_id);
    setSelectedDistrict({
      id: school.district_id,
      name: school.district
    });
  }

  const closeEditModal = () => {
    // First start the closing animation
    setIsEditModalAnimating(true);
    setFormData({
      id: 0,
      code: '',
      name: '',
      office_id: 0,
      is_public_school: false,
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

  const [formData, setFormData] = useState<CreateUpdateSchool>({
    id: 0,
    code: '',
    name: '',
    office_id: 0,
    is_public_school: true,
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

  // Fetch schools from API
  const fetchSchools = async (page: number = 1) => {
    setLoading(true);
    setTimeout(async() => {
      try {
        const response = await getSchools(page, metadata.per_page);
        
        setSchools(response.data);
        setMetadata(response.meta_data);
      } catch (error) {
        console.error('Error fetching schools:', error);
      } finally {
        setLoading(false);
      }
    }, 500);
  };

  useEffect(() => {
    fetchSchools();
  }, []);

  

  // ------------------------------ LOCATION HANDLER -----------------------------
  // Fetch provinces from API
  const [provinces, setProvinces] = useState<Province[]>([]);
  const [loadingProvinces, setLoadingProvinces] = useState<boolean>(false);
  const [selectedProvince, setSelectedProvince] = useState<Province | null>(null);
  const [provinceSearchTerm, setProvinceSearchTerm] = useState<string>('');
  const filteredProvinces = provinceSearchTerm === '' 
    ? provinces 
    : provinces.filter((province) => 
        province.name.toLowerCase().includes(provinceSearchTerm.toLowerCase())
      );

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

  // Fetch regencies from API
  const [regencies, setRegencies] = useState<Regency[]>([]);
  const [loadingRegencies, setLoadingRegencies] = useState<boolean>(false);
  const [selectedRegency, setSelectedRegency] = useState<Regency | null>(null);
  const [regencySearchTerm, setRegencySearchTerm] = useState<string>('');
  const filteredRegencies = regencySearchTerm === '' 
    ? regencies 
    : regencies.filter((regency) => 
        regency.name.toLowerCase().includes(regencySearchTerm.toLowerCase())
      );

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

  // Fetch districts from API
  const [districts, setDistricts] = useState<District[]>([]);
  const [loadingDistricts, setLoadingDistricts] = useState<boolean>(false);
  const [selectedDistrict, setSelectedDistrict] = useState<District | null>(null);
  const [districtSearchTerm, setDistrictSearchTerm] = useState<string>('');
  const filteredDistricts = districtSearchTerm === '' 
    ? districts 
    : districts.filter((district) => 
        district.name.toLowerCase().includes(districtSearchTerm.toLowerCase())
      );
  
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

  // ------------------------------ OFFICE HANDLER -----------------------------
  // Fetch office from API
  const [loadingOffices, setLoadingOffices] = useState<boolean>(false);
  const [selectedOffice, setSelectedOffice] = useState<SelectOption | null>(null);
  const [officeSearchTerm, setOfficeSearchTerm] = useState<string>('');
  const [offices, setOffices] = useState<SelectOption[]>([]);
  const filteredOffices = officeSearchTerm === '' 
     ? offices 
     : offices.filter((office) => 
         office.name.toLowerCase().includes(officeSearchTerm.toLowerCase())
       );

  const fetchOffices = async (searchTerm: string = '') => {
    setLoadingOffices(true);
    setTimeout(async() => {
      try {
        const response = await fetchOfficesSelectOption(searchTerm);
        if (response.data.length > 0) {
          setOffices(response.data);
        }
      } catch (error) {
        console.error('Error fetching offices:', error);
      } finally {
        setLoadingOffices(false);
      }
    }, 500);
  };

  const handleOfficeQueryChange = (query: string) => {
    setOfficeSearchTerm(query);
    fetchOffices(query);
  };

  const handleOfficeChange = (office: SelectOption) => {
    fetchOffices(office.id.toString());
    setSelectedOffice(office);
    setSelectedProvince(office.additional_data?.province);
    setSelectedRegency(office.additional_data?.regency);
    setSelectedDistrict(office.additional_data?.district);
    setFormData(prev => ({
      ...prev,
      office_id: office.id,
      office: office.name,
      province_id: office.additional_data?.province?.id,
      province: office.additional_data?.province?.name,
      regency_id: office.additional_data?.regency?.id,
      regency: office.additional_data?.regency?.name,
      district_id: office.additional_data?.district?.id,
      district: office.additional_data?.district?.name
    }));
    
    // Clear error for province field
    if (formErrors.office) {
      setFormErrors(prev => {
        const newErrors = {...prev};
        delete newErrors.office;
        return newErrors;
      });
    }
  };

  const officeData = {
      data: {
        values: offices,
        label: "Kantor",
        isLoading: loadingOffices,
        selectedValue: selectedOffice,
        searchTerm: officeSearchTerm,
        filteredOptions: filteredOffices,
        fieldError: formErrors.office_id,
        placeholder: "Pilih Kantor"
      },
      onChange: handleOfficeChange,
      handleQueryChange: handleOfficeQueryChange,
      handleClearSearch: () => {
        setOfficeSearchTerm('');
        fetchOffices();
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
    console.log(formData);
    const errors: {[key: string]: string} = {};
    
    if (!formData.code.trim()) errors.code = "Kode sekolah wajib diisi";
    if (!formData.name.trim()) errors.name = "Nama sekolah wajib diisi";
    if (!formData.office_id) errors.office_id = "Kantor wajib diisi";
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

  // Submit new school
  const handleSubmit = async () => {
    if (!validateForm()) return;
    
    setIsSubmitting(true);
    try {
      const response = await createSchool(formData);
      
      // Add the new school to the list and close modal
      setSchools(prev => [response.data, ...prev]);
      closeModal();
      
      // Refresh the list to ensure correct pagination
      fetchSchools(metadata.current_page);
    } catch (error) {
      console.error('Error creating school:', error);
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

  // Submit new school
  const handleEditSubmit = async () => {
    if (!validateForm()) return;
    
    setIsSubmitting(true);
    try {
      if (!formData.id) {
        setIsSubmitting(false);
        console.error('School ID is required for update');
        return;
      }
      const response = await updateSchool(formData.id, formData);
      
      // Add the new school to the list and close modal
      setSchools(prev => [response.data, ...prev]);
      closeEditModal();
      
      // Refresh the list to ensure correct pagination
      fetchSchools(metadata.current_page);
    } catch (error) {
      console.error('Error creating school:', error);
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
      fetchSchools(newPage);
    }
  };

  const filteredSchools = schools.filter(school => 
    school.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    school.code.toLowerCase().includes(searchTerm.toLowerCase()) ||
    school.province.toLowerCase().includes(searchTerm.toLowerCase()) ||
    school.regency.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const totalPages = Math.ceil(metadata.total_items / metadata.per_page);

  // Open delete confirmation modal
  const openDeleteModal = (school: SchoolType) => {
    setSchoolToDelete({
      id: school.id,
      label: moduleName,
      name: school.name
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
      setSchoolToDelete(null);
    }, 300);
  };

  // Handle delete school
  const handleDelete = async () => {
    if (!schoolToDelete) return;
    setIsDeleting(true);
    
    try {
      await deleteSchool(Number(schoolToDelete.id));
      // Refresh the schools list
      fetchSchools(metadata.current_page);
      closeDeleteModal();
    } catch (error) {
      console.error('Failed to delete school:', error);
    } finally {
      setIsDeleting(false);
    }
  };

  return (
    <div className="container mx-auto">
      <div className="flex justify-between items-center mb-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Sekolah</h1>
          <p className="mt-1 text-sm text-gray-500">
          Kelola semua sekolah dalam sistem
          </p>
        </div>
        <button 
          className="bg-amber-500 hover:bg-amber-600 text-white px-4 py-2 rounded-md flex items-center"
          onClick={openModal}
        >
          <FiPlus className="mr-2" />
          Tambah Sekolah
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
              placeholder="Cari sekolah..."
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
                <th scope="col" className="w-16 px-3 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  #
                </th>
                <th scope="col" className="w-24 px-3 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Kode
                </th>
                <th scope="col" className="w-48 px-3 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Nama
                </th>
                <th scope="col" className="w-48 px-3 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Kantor
                </th>
                <th scope="col" className="w-64 px-3 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Lokasi
                </th>
                <th scope="col" className="w-32 px-3 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Kepemilikan
                </th>
                <th scope="col" className="w-48 px-3 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Kontak
                </th>
                <th scope="col" className="w-24 px-3 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Aksi
                </th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {loading ? (
                <tr>
                  <td colSpan={8} className="px-6 py-4 text-center">
                    <div className="flex justify-center">
                      <div className="animate-spin rounded-full h-6 w-6 border-b-2 border-amber-500"></div>
                    </div>
                  </td>
                </tr>
              ) : filteredSchools.length === 0 ? (
                <tr>
                  <td colSpan={8} className="px-6 py-4 text-center text-sm text-gray-500">
                    Tidak ada sekolah...
                  </td>
                </tr>
              ) : (
                filteredSchools.map((school, index) => (
                  <tr key={school.id} className="hover:bg-gray-50">
                    <td className="px-3 py-3 whitespace-nowrap text-sm font-medium text-gray-900">
                      {((metadata.current_page - 1) * metadata.per_page) + index + 1}
                    </td>
                    <td className="px-3 py-3 whitespace-nowrap text-sm font-medium text-gray-900">
                      {school.code}
                    </td>
                    <td className="px-3 py-3 text-sm text-gray-900 truncate max-w-[12rem]" title={school.name}>
                      {school.name}
                    </td>
                    <td className="px-3 py-3 text-sm text-gray-500 truncate max-w-[12rem]" title={school.office}>
                      {school.office}
                    </td>
                    <td className="px-3 py-3 text-sm text-gray-500">
                      <div className="flex flex-col space-y-0.5">
                        <span className="truncate max-w-[14rem]" title={school.province}>
                          {school.province}
                        </span>
                        <span className="text-xs text-gray-400 truncate max-w-[14rem]" title={`${school.regency}, ${school.district}`}>
                          {school.regency}, {school.district}
                        </span>
                      </div>
                    </td>
                    <td className="px-3 py-3 whitespace-nowrap text-center">
                      <span className="px-2 py-1 inline-flex text-xs leading-4 font-semibold rounded-full bg-amber-100 text-amber-800">
                        {school.is_public_school ? "Pemerintah" : "Swasta"}
                      </span>
                    </td>
                    <td className="px-3 py-3 text-sm text-gray-500">
                      <div className="flex flex-col space-y-0.5">
                        <span className="truncate max-w-[11rem]" title={school.email}>
                          {school.email}
                        </span>
                        {school.phone && (
                          <span className="text-xs text-gray-400 truncate max-w-[11rem]" title={school.phone}>
                            {school.phone}
                          </span>
                        )}
                      </div>
                    </td>
                    <td className="px-3 py-3 whitespace-nowrap text-right text-sm font-medium">
                      <div className="flex justify-end space-x-1">
                        <button 
                          className="text-amber-600 hover:text-amber-900"
                          onClick={() => openEditModal(school)}
                        >
                          <div className="flex items-center justify-center h-8 w-8 rounded-full bg-amber-100 hover:bg-amber-200">
                            <FiEdit size={16} />
                          </div>
                        </button>
                        <button 
                          className="text-red-600 hover:text-red-900"
                          onClick={() => openDeleteModal(school)}
                        >
                          <div className="flex items-center justify-center h-8 w-8 rounded-full bg-red-100 hover:bg-red-200">
                            <FiTrash2 size={16} />
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
        <MasterSchoolModal
          formTitle="Tambah Sekolah Baru"
          isSubmitting={isSubmitting}
          isModalAnimating={isModalAnimating}
          formData={formData}
          formErrors={formErrors}
          officeData={officeData}
          locationData={locationData}
          closeModal={closeModal}
          onAnimationEnd={() => setIsModalAnimating(false)}
          handleInputChange={handleInputChange}
          handleSubmit={handleSubmit}
        />
      )}

      {isEditModalOpen && (
        <MasterSchoolModal
          formTitle="Ubah Sekolah"
          isSubmitting={isSubmitting}
          isModalAnimating={isEditModalAnimating}
          formData={formData}
          formErrors={formErrors}
          officeData={officeData}
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
          deleteModalData={schoolToDelete}
          handleDelete={handleDelete}
          closeDeleteModal={closeDeleteModal}
          onAnimationEnd={() => setIsDeleteModalAnimating(false)}
        />
      )}
    </div>
  );
};

export default School;