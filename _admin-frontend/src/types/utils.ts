
export interface SelectOption {
  id: number;
  name: string;
  additional_data?: any | null;
}

export interface SelectOptionData {
  label: string;
  values: SelectOption[];
  isLoading: boolean;
  selectedValue: SelectOption | null;
  searchTerm: string;
  filteredOptions: SelectOption[];
  fieldError: string | null;
  placeholder?: string;
}

export interface SelectOptionProps {
  data: SelectOptionData;
  onChange: (value: SelectOption) => void;
  handleQueryChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  handleClearSearch: () => void;
}