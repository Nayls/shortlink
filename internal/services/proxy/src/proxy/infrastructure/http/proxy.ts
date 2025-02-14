import * as express from 'express'
import { inject } from 'inversify'
import { interfaces, controller, httpGet, httpPost, request, response } from "inversify-express-utils"
import { Logger } from "tslog"

import { LinkService } from '../../service/links'
import TYPES from '../../../types'
import {StatsService} from "../../service/stats";

const log: Logger = new Logger()

@controller(`/s/:hash`)
class ProxyController implements interfaces.Controller {
  constructor(
    @inject(TYPES.SERVICE.LinkService) private linkService: LinkService,
    @inject(TYPES.SERVICE.StatsService) private statsService: StatsService,
  ) { }

  @httpGet(`/`)
  public async redirect (@request() req: express.Request, @response() res: express.Response) {
    const { hash } = req.params

    try {
      // TODO: request by gRPC to LinkService: getLinkByHash
      const link = await this.linkService.get(hash)
      res.redirect(301, link)
    } catch (error) {
      res.status(400).json(error)
    }
  }
}

export default ProxyController
