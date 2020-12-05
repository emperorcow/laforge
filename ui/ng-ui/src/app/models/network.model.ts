import { ID, varsMap, Tag, Status } from './common.model';
import { Build } from './environment.model';
import { ProvisionedHost } from './host.model';

interface Network {
  id: ID;
  name: String;
  cidr: String;
  vdiVisible: Boolean;
  vars: varsMap[];
  tags: Tag[];
}

interface ProvisionedNetwork {
  id: ID;
  name: String;
  cidr: String;
  vars: varsMap[];
  tags: Tag[];
  provisionedHosts: ProvisionedHost[];
  status: Status;
  network: Network;
  build: Build;
}

export {
  Network,
  ProvisionedNetwork
}