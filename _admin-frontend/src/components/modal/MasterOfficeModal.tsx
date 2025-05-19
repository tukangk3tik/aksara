import { FiPlus, FiX } from "react-icons/fi";
import Button from "../Button";
import SelectSearchOption from "../SelectOption";
import { CreateUpdateOffice } from "../../types/office";
import { SelectOptionProps } from "../../types/utils";
import React from "react";

export interface LocationData {
  province: SelectOptionProps;
  regency: SelectOptionProps;
  district: SelectOptionProps;
}

interface MasterOfficeModalProps {
  formTitle: string;
  isSubmitting: boolean;
  isModalAnimating: boolean;
  formData: CreateUpdateOffice;
  formErrors: {[key: string]: string};
  locationData: LocationData;
  closeModal: () => void;
  onAnimationEnd: () => void;
  handleInputChange: (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => void;
  handleSubmit: () => void;
}

const MasterOfficeModal: React.FC<MasterOfficeModalProps> = ({
  formTitle,
  isSubmitting,
  isModalAnimating,
  formData,
  formErrors,
  locationData,
  closeModal,
  onAnimationEnd,
  handleInputChange,
  handleSubmit
}) => {
  const isEdit = formData.id ? true : false;

  return (
    <>
    {/* Backdrop */}
    <div className="fixed inset-0 bg-gray-500/75 transition-opacity duration-300 z-40" 
         aria-hidden="true"
         onClick={closeModal} />

    {/* Modal container - centered with flex */}
    <div className="fixed inset-0 z-50 overflow-y-auto">
      <div className="flex min-h-full items-center justify-center p-4 text-center sm:p-0">

        {/* Modal panel with animation */}
        <div className={`relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-xl ${isModalAnimating ? 'animate-modal-out' : 'animate-modal-in'}`}
            onAnimationEnd={onAnimationEnd}>
          {/* Modal header */}
          <div className="border-b border-gray-200 px-4 py-3 sm:px-6 flex items-center justify-between">
            <h3 className="text-base font-semibold leading-6 text-gray-900 flex items-center">
              <FiPlus className="mr-2 text-amber-500" />
              {formTitle} 
            </h3>
            <button 
              type="button" 
              className="text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-amber-500 focus:ring-offset-2 rounded-md p-1.5"
              onClick={closeModal}
            >
              <span className="sr-only">Close</span>
              <FiX className="h-5 w-5" />
            </button>
          </div>
           
          {/* Modal body */}
          <div className="bg-white px-4 py-4 sm:p-6">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {/* Your existing form fields - keeping them intact */}
              {/* Code */}
              <div>
                <label htmlFor="code" className="block text-sm font-medium leading-6 text-gray-900">
                  Kode Kantor <span className="text-red-500">*</span>
                </label>
                <div className="mt-1">
                  <input
                    disabled={isEdit}
                    type="text"
                    id="code"
                    name="code"
                    placeholder="Contoh: KTR001"
                    value={formData.code}
                    onChange={handleInputChange}
                    className={`block w-full rounded-md px-2 py-1.5 text-sm shadow-sm ring-1 
                      ${formErrors.code ? 'ring-red-500' : 'ring-gray-300'} focus:ring-2 focus:ring-amber-500 
                      ${isEdit ? 'bg-gray-200' : 'bg-white'}`}
                  />
                  {formErrors.code && (
                    <p className="mt-1 text-xs text-red-500 flex items-center">
                      <span className="mr-1">⚠</span> {formErrors.code}
                    </p>
                  )}
                </div>
              </div>
                
              {/* Name */}
              <div>
                <label htmlFor="name" className="block text-sm font-medium leading-6 text-gray-900">
                  Nama Kantor <span className="text-red-500">*</span>
                </label>
                <div className="mt-1">
                  <input
                    type="text"
                    id="name"
                    name="name"
                    placeholder="Nama kantor"
                    value={formData.name}
                    onChange={handleInputChange}
                    className={`block w-full rounded-md px-2 py-1.5 text-sm shadow-sm ring-1 ${formErrors.name ? 'ring-red-500' : 'ring-gray-300'} focus:ring-2 focus:ring-amber-500`}
                  />
                  {formErrors.name && (
                    <p className="mt-1 text-xs text-red-500 flex items-center">
                      <span className="mr-1">⚠</span> {formErrors.name}
                    </p>
                  )}
                </div>
              </div>
                
              {/* Province */}
              <SelectSearchOption
                data={{
                  values: locationData.province.data.values,
                  label: "Provinsi",
                  isLoading: locationData.province.data.isLoading,
                  selectedValue: locationData.province.data.selectedValue,
                  searchTerm: locationData.province.data.searchTerm,
                  filteredOptions: locationData.province.data.filteredOptions,
                  fieldError: formErrors.province_id,
                  placeholder: "Pilih Provinsi"
                }}
                isDisable={isEdit}
                handleChange={locationData.province.onChange}
                handleQueryChange={locationData.province.handleQueryChange}
                handleClearSearch={() => {
                  locationData.province.handleClearSearch();
                }}
              />

              {/* Regency */}
              <SelectSearchOption
                data={{
                  values: locationData.regency.data.values,
                  label: "Kabupaten/Kota",
                  isLoading: locationData.regency.data.isLoading,
                  selectedValue: locationData.regency.data.selectedValue,
                  searchTerm: locationData.regency.data.searchTerm,
                  filteredOptions: locationData.regency.data.filteredOptions,
                  fieldError: formErrors.regency_id,
                  placeholder: "Pilih Kabupaten/Kota"
                }}
                isDisable={isEdit}
                handleChange={locationData.regency.onChange}
                handleQueryChange={locationData.regency.handleQueryChange}
                handleClearSearch={() => {
                  locationData.regency.handleClearSearch();
                }}
              />
             
              {/* District */}
              <SelectSearchOption
                data={{
                  values: locationData.district.data.values,
                  label: "Kecamatan",
                  isLoading: locationData.district.data.isLoading,
                  selectedValue: locationData.district.data.selectedValue,
                  searchTerm: locationData.district.data.searchTerm,
                  filteredOptions: locationData.district.data.filteredOptions,
                  fieldError: formErrors.district_id,
                  placeholder: "Pilih Kecamatan"
                }}
                isDisable={isEdit}
                handleChange={locationData.district.onChange}
                handleQueryChange={locationData.district.handleQueryChange}
                handleClearSearch={() => {
                  locationData.district.handleClearSearch();
                }}
              />
             
              {/* Email */}
              <div>
                <label htmlFor="email" className="block text-sm font-medium leading-6 text-gray-900">
                  Email
                </label>
                <div className="mt-1">
                  <input
                    type="email"
                    id="email"
                    name="email"
                    placeholder="email@example.com"
                    value={formData.email}
                    onChange={handleInputChange}
                    className={`block w-full rounded-md px-2 py-1.5 text-sm shadow-sm ring-1 ${formErrors.email ? 'ring-red-500' : 'ring-gray-300'} focus:ring-2 focus:ring-amber-500`}
                  />
                  {formErrors.email && (
                    <p className="mt-1 text-xs text-red-500 flex items-center">
                      <span className="mr-1">⚠</span> {formErrors.email}
                    </p>
                  )}
                </div>
              </div>
                
              {/* Phone */}
              <div>
                <label htmlFor="phone" className="block text-sm font-medium leading-6 text-gray-900">
                  Nomor Telepon
                </label>
                <div className="mt-1">
                  <input
                    type="text"
                    id="phone"
                    name="phone"
                    placeholder="+62..."
                    value={formData.phone}
                    onChange={handleInputChange}
                    className="block w-full rounded-md px-2 py-1.5 text-sm shadow-sm ring-1 ring-gray-300 focus:ring-2 focus:ring-amber-500"
                  />
                </div>
              </div>
                
              {/* Address - full width */}
              <div className="md:col-span-2">
                <label htmlFor="address" className="block text-sm font-medium leading-6 text-gray-900">
                  Alamat
                </label>
                <div className="mt-1">
                  <textarea
                    id="address"
                    name="address"
                    rows={2}
                    placeholder="Alamat lengkap"
                    value={formData.address}
                    onChange={handleInputChange}
                    className="block w-full rounded-md px-2 py-1.5 text-sm shadow-sm ring-1 ring-gray-300 focus:ring-2 focus:ring-amber-500"
                  />
                </div>
              </div>
            </div>
          </div>
                
          {/* Modal footer */}
          <div className="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
            <Button
              type="button"
              variant="primary"
              size="md"
              isLoading={isSubmitting}
              icon={<FiPlus className="h-3.5 w-3.5" />}
              onClick={handleSubmit}
              className="sm:ml-3 sm:w-auto"
            >
              {isSubmitting ? 'Menyimpan...' : 'Simpan'}
            </Button>
            <Button 
              type="button"
              variant="secondary"
              size="md"
              onClick={closeModal}
              disabled={isSubmitting}
            >
              Batal
            </Button>
          </div>
        </div>
      </div>
    </div>
  </>
  );
};

export default MasterOfficeModal;