// types.ts
export type Data = {
  id: string;
  name: string;
  age: number;
  occupation: string;
  address: {
    street: string;
    city: string;
    state: string;
    zip: string;
  };
};

export type DataArray = Data[];

export type GroupByField = 'id' | 'name' | 'age' | 'occupation' | 'address.street' | 'address.city' | 'address.state' | 'address.zip';

export type GroupByOptions = {
  field: GroupByField;
  sort?: 'asc' | 'desc';
};