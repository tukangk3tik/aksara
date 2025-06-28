import { Listbox, ListboxButton, ListboxOption, ListboxOptions, Transition } from "@headlessui/react";
import { FiChevronDown, FiSearch, FiX } from "react-icons/fi";
import { Fragment } from "react";
import { SelectOption, SelectOptionData } from "../types/utils";


interface SelectSearchOptionProps {
  data: SelectOptionData;
  isDisable: boolean;
  handleChange: (value: SelectOption) => void;
  handleQueryChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  handleClearSearch: (e?: React.MouseEvent) => void;
}

const SelectSearchOption: React.FC<SelectSearchOptionProps> = ({ 
  data,
  isDisable = false,
  handleChange,
  handleQueryChange,
  handleClearSearch,
}) => {
  return (
    <div>
    <label className="block text-sm font-medium leading-6 text-gray-900">
      {data.label} <span className="text-red-500">*</span>
    </label>
    <div className="mt-1 relative">
      <Listbox value={data.selectedValue} onChange={handleChange} disabled={isDisable}>
        <div className="relative">
          <ListboxButton className={`relative w-full cursor-default rounded-md py-1.5 pl-3 pr-10 text-left text-sm shadow-sm ring-1 
            ${data.fieldError ? 'ring-red-500' : 'ring-gray-300'} focus:ring-2 focus:ring-amber-500 h-9
            ${isDisable ? 'bg-gray-200' : 'bg-white'}`}
          >
            {data.isLoading ? (
              <span className="flex items-center text-gray-500">
                <div className="animate-spin rounded-full h-3 w-3 border-b-2 border-gray-500 mr-2"></div>
                Memuat...
              </span>
            ) : (
              <>
                <span className="block truncate">
                  {data.selectedValue ? data.selectedValue.name : data.placeholder}
                </span>
                <span className="absolute inset-y-0 right-0 flex items-center pr-2">
                  <FiChevronDown className="h-4 w-4 text-gray-400" aria-hidden="true" />
                </span>
              </>
            )}
          </ListboxButton>
          <Transition
            as={Fragment}
            leave="transition ease-in duration-100"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <ListboxOptions className="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
              <div className="sticky top-0 bg-white px-3 py-2 border-b border-gray-200 z-10">
                <div className="relative">
                  <FiSearch className="absolute left-2 top-1/2 -translate-y-1/2 text-gray-400" />
                  <input
                    type="text"
                    className="w-full pl-8 pr-8 py-1.5 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-amber-500"
                    placeholder={`Cari ${data.label.toLowerCase()}...`}
                    value={data.searchTerm}
                    onChange={handleQueryChange}
                    onClick={(e) => e.stopPropagation()}
                    onKeyDown={(e) => {
                      if (e.key === 'Enter') e.preventDefault();
                    }}
                  />
                  {data.searchTerm && (
                    <button
                      type="button"
                      className="absolute right-2 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
                      onClick={(e) => {
                        e.stopPropagation();
                        handleClearSearch(e);
                      }}
                    >
                      <FiX className="h-3.5 w-3.5" />
                    </button>
                  )}
                </div>
              </div>
              
              <div className="max-h-48 overflow-y-auto mt-1">
                {data.filteredOptions.length === 0 ? (
                  <div className="px-3 py-2 text-gray-500">
                    {data.values.length === 0 ? 'Tidak ada data' : 'Tidak ada hasil pencarian'}
                  </div>
                ) : (
                  data.filteredOptions.map((item) => (
                    <ListboxOption
                      key={item.id}
                      value={item}
                      className={({ selected }) =>
                        `relative cursor-default select-none py-2 pl-3 pr-9 ${
                          selected ? 'bg-amber-100 text-amber-900' : 'text-gray-900'
                        }`
                      }
                    >
                      {({ selected }) => (
                        <>
                          <span className={`block truncate ${selected ? 'font-semibold' : 'font-normal'}`}>
                            {item.name}
                          </span>
                        </>
                      )}
                    </ListboxOption>
                  ))
                )}
              </div>
            </ListboxOptions>
          </Transition>
        </div>
      </Listbox>
      {data.fieldError && (
        <p className="mt-1 text-xs text-red-500 flex items-center">
          <span className="mr-1">⚠</span> {data.fieldError}
        </p>
      )}
    </div>
  </div>
  );
};

export default SelectSearchOption;