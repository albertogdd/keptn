import { Sequence } from './sequence';
import { Approval } from '../../shared/interfaces/approval';
import { DeploymentInformation, Service as sv } from '../../shared/models/service';
import { Remediation } from '../../shared/models/remediation';
import { EventTypes } from '../../shared/interfaces/event-types';
import { IServiceEvent } from '../../shared/interfaces/service';

export class Service extends sv {
  lastEventTypes: { [event: string]: IServiceEvent | undefined } = {};
  latestSequence?: Sequence;
  openRemediations: Remediation[] = [];
  openApprovals: Approval[] = [];
  deploymentInformation?: DeploymentInformation;

  public static fromJSON(data: unknown): Service {
    return Object.assign(new this(), data);
  }

  public get latestDeploymentEvent(): IServiceEvent | undefined {
    return this.deploymentEvent ?? this.evaluationEvent;
  }

  public get deploymentEvent(): IServiceEvent | undefined {
    return this.lastEventTypes?.[EventTypes.DEPLOYMENT_FINISHED];
  }

  private get evaluationEvent(): IServiceEvent | undefined {
    return this.lastEventTypes?.[EventTypes.EVALUATION_FINISHED];
  }
}
